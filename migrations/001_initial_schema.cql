CREATE KEYSPACE IF NOT EXISTS bento_scheduler
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  } AND DURABLE_WRITES = true;

USE bento_scheduler;

CREATE TABLE IF NOT EXISTS scheduled_jobs (
  id UUID,
  job_type TEXT,
  callback_url TEXT,
  payload TEXT,
  scheduled_time TIMESTAMP,
  status TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  execution_time TIMESTAMP,
  retry_count INT,
  error TEXT,
  PRIMARY KEY ((status), scheduled_time, id)
) WITH CLUSTERING ORDER BY (scheduled_time ASC, id ASC)
  AND gc_grace_seconds = 864000;