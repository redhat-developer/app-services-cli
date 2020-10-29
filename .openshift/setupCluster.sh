oc new-project mas-mock > /dev/null

oc apply -f ./kafka.yml
oc apply -f ./keycloak.yml