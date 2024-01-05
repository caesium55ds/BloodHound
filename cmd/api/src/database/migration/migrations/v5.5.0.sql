ALTER TABLE audit_logs 
ADD COLUMN IF NOT EXISTS actor_email VARCHAR(330) DEFAULT NULL,
ADD COLUMN IF NOT EXISTS source VARCHAR(40) DEFAULT NULL,
ADD COLUMN IF NOT EXISTS status VARCHAR(15) CHECK (status IN ('success', 'failure')) DEFAULT 'success';


