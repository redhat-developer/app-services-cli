var topics = require('../_data_/topics.json')

module.exports = {
  createTopic: async (c, req, res) => {
    const topicBody = c.request.body;

    if (!topicBody) {
      return res.status(400).json({ err: 'Bad request' })
    }

    let topic = getTopic(topicBody.name)

    if (topic) {
      return res.status(409).json({ err: "topic exists" })
    }

    topic = {
      name: topicBody.name,
      config: topicBody.settings.config,
      partitions: createPartitions(topicBody.settings.numPartitions, topicBody.settings.replicationFactor)
    }
    topics.push(topic)
    return res.status(200).json(topic)
  },

  getTopicsList: async (c, req, res) => {
    return res.status(200).json({
      limit: parseInt(req.query.limit, 10) || 100,
      offset: 0,
      count: topics?.length,
      topics: topics
    });
  },

  getTopic: async (c, req, res) => {
    const topic = getTopic(c.request.params.topicName)
    if (!topic) {
      return res.status(404).json({ err: "not found" })
    }
    return res.status(200).json(topic)
  },

  deleteTopic: async (c, _, res) => {
    const topicName = c.request.params.topicName;

    const topic = getTopic(topicName)
    if (!topic) {
      return res.status(404).json({ err: "not found" })
    }
    topics = topics.filter(t => t.name !== topicName);

    return res.status(200).json({ message: 'deleted' })
  },

  updateTopic: async (c, _, res) => {
    const topicName = c.request.params.topicName;
    const topicBody = c.request.body;

    if (!topicBody) {
      return res.status(400).json({ err: 'Bad request' })
    }

    const topic = getTopic(topicName)
    if (topicBody.numPartitions) {
      topic.partitions = createPartitions(topicBody.numPartitions, 2)
    }
    if (topicBody.config) {
      topic.config = topicBody.config;
    }

    if (!topic) {
      return res.status(404).json({ err: "not found" })
    }

    if (topicBody?.settings?.config) {
      topic.config = topicBody.settings.config;
    }

    return res.status(200).json(topic)
  },

  // Handling auth
  notFound: async (c, req, res) => {
    debug(res)
    return res.status(404).json({ err: "not found" })
  },
  unauthorizedHandler: async (c, req, res) =>
    res.status(401).json({ err: "unauthorized" }),
};

function getTopic(name) {
  const topic = topics.find(t => t.name === name);

  return topic
}

const createPartitions = (numberOfPartitions, numberOfReplicas) => {
  const partitions = []
  for (let i = 0; i < numberOfPartitions; i++) {
    const id = i + 1;
    partitions.push(createPartition(id, numberOfReplicas))
  }

  return partitions
}

const createPartition = (id, numberOfReplicas) => {
  const replicas = []
  for (let i = 0; i < numberOfReplicas; i++) {
    replicas.push({ id: i + 1 })
  }
  return {
    id,
    leader: {
      id: 1
    },
    replicas
  }
}

