/**
 * This script helps set up a local keycloak environment
 * It assumes keycloak is already running and available at
 * http://localhost:8080
 * It does the following:
 * * Imports a new realm located in ./realm-export.json. This realm has two clients
 *   * datasync-starter-client- public client used by the mobile application
 *   * datasync-starter-server - bearer client used by the server to authenticate requests
 * * Creates the realm roles and client roles defined in the realmRoleNames and clientRoleNames lists
 * * Creates the users defined in the users list
 * * Assigns the client and realm roles to the users
 */

const axios = require("axios");
const realmToImport = require("./realm-export.json");
const { writeFileSync } = require("fs");

// the keycloak server we're working against
const KEYCLOAK_URL = process.env.KEYCLOAK_URL || "http://localhost:8080/auth";

// name of the realm
const APP_REALM = process.env.KEYCLOAK_REALM || "sso-external";

// name of the admin realm
const ADMIN_REALM = "master";

const RESOURCE = "admin-cli";
const ADMIN_USERNAME = process.env.KEYCLOAK_ADMIN_USERNAME || "admin";
const ADMIN_PASSWORD = process.env.KEYCLOAK_ADMIN_PASSWORD || "admin";
let token = "";

// The keycloak client used by the sample app
const PUBLIC_CLIENT_NAME = "rhoas-cli-prod";
let PUBLIC_CLIENT;

// The client roles you want created for the BEARER_CLIENT_NAME client
const clientRoleNames = [];

// The realm roles we want for the realm
const realmRoleNames = [];

let realmRoles;

const writeConfig = false;

// This is called by an immediately invoked function expression
// at the bottom of the file
async function prepareKeycloak() {
  try {
    console.log("Authenticating with keycloak server");
    token = await authenticateKeycloak();

    // Always do a hard reset first just to keep things tidy
    console.log("Going to reset keycloak");
    await resetKeycloakConfiguration();

    console.log("Importing sample realm into keycloak");
    await importRealm();

    console.log("Fetching available clients from keycloak");
    const clients = await getClients();

    // Get the public client object from keycloak
    // Need this for the ID assigned by keycloak
    PUBLIC_CLIENT = clients.find(
      (client) => client.clientId === PUBLIC_CLIENT_NAME
    );

    console.log("creating client roles");
    for (let roleName of clientRoleNames) {
      await createClientRole(PUBLIC_CLIENT, roleName);
    }

    console.log("creating realm roles");
    for (let roleName of realmRoleNames) {
      await createRealmRole(roleName);
    }

    // get the actual role objects from keycloak after creating them
    // need to get the ids that were created on them
    realmRoles = await getRealmRoles();
    publicClientRoles = await getClientRoles(PUBLIC_CLIENT);

    const publicInstallation = await getClientInstallation(PUBLIC_CLIENT);

    if (writeConfig) {
      writeFileSync(
        `../client/public/keycloak.json`,
        JSON.stringify(publicInstallation, null, 2)
      );
    }
    console.log();
    console.log(
      "Your keycloak server is set up for local usage and development"
    );
  } catch (e) {
    console.error(e);
    process.exit(1);
  }
}

async function getClientInstallation(
  client,
  installationType = "keycloak-oidc-keycloak-json"
) {
  if (client) {
    const res = await axios({
      method: "GET",
      url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/clients/${client.id}/installation/providers/${installationType}`,
      headers: { Authorization: token },
    });
    return res.data;
  }
  throw new Error("client is undefined");
}

async function authenticateKeycloak() {
  const res = await axios({
    method: "POST",
    url: `${KEYCLOAK_URL}/realms/${ADMIN_REALM}/protocol/openid-connect/token`,
    data: `client_id=${RESOURCE}&username=${ADMIN_USERNAME}&password=${ADMIN_PASSWORD}&grant_type=password`,
  });
  return `Bearer ${res.data["access_token"]}`;
}

async function importRealm() {
  return await axios({
    method: "POST",
    url: `${KEYCLOAK_URL}/admin/realms`,
    data: realmToImport,
    headers: { Authorization: token, "Content-Type": "application/json" },
  });
}

async function getRealmRoles() {
  const res = await axios({
    method: "GET",
    url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/roles`,
    headers: { Authorization: token },
  });
  return res.data;
}

async function createClientRole(client, roleName) {
  try {
    return await axios({
      method: "POST",
      url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/clients/${client.id}/roles`,
      headers: { Authorization: token },
      data: {
        clientRole: true,
        name: roleName,
      },
    });
  } catch (e) {
    if (
      e.response.data.errorMessage ===
      `Role with name ${roleName} already exists`
    ) {
      console.log(e.response.data.errorMessage);
    } else {
      throw e;
    }
  }
}

async function createRealmRole(roleName) {
  try {
    return await axios({
      method: "POST",
      url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/roles`,
      headers: { Authorization: token },
      data: {
        clientRole: false,
        name: roleName,
      },
    });
  } catch (e) {
    if (
      e.response.data.errorMessage ===
      `Role with name ${roleName} already exists`
    ) {
      console.log(e.response.data.errorMessage);
    } else {
      throw e;
    }
  }
}

async function getClients() {
  const res = await axios({
    method: "GET",
    url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/clients`,
    headers: { Authorization: token },
  });
  return res.data;
}

async function getClientRoles(client) {
  const res = await axios({
    method: "GET",
    url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}/clients/${client.id}/roles`,
    headers: { Authorization: token },
  });
  return res.data;
}

async function assignRealmRoleToUser(userIdUrl, role) {
  const res = await axios({
    method: "POST",
    url: `${userIdUrl}/role-mappings/realm`,
    data: [role],
    headers: { Authorization: token, "Content-Type": "application/json" },
  });
  return res.data;
}

async function assignClientRoleToUser(userIdUrl, client, role) {
  const res = await axios({
    method: "POST",
    url: `${userIdUrl}/role-mappings/clients/${client.id}`,
    data: [role],
    headers: { Authorization: token, "Content-Type": "application/json" },
  });
  return res.data;
}

async function resetKeycloakConfiguration() {
  try {
    await axios({
      method: "DELETE",
      url: `${KEYCLOAK_URL}/admin/realms/${APP_REALM}`,
      headers: { Authorization: token },
    });
  } catch (e) {
    if (e.response.status !== 404) {
      throw e;
    }
    console.log(`404 while deleting realm ${APP_REALM} - ignoring`);
  }
}

(async () => {
  await prepareKeycloak();
})();
