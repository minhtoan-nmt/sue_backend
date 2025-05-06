
-- ============================================
-- 1. CREATE TABLES (WITHOUT FOREIGN KEYS FIRST)
-- ============================================

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    salt VARCHAR(50) NOT NULL,
    phone VARCHAR(20),
    role VARCHAR(20) CHECK (role IN ('Admin', 'Teacher', 'Student')) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('active', 'inactive', 'deleted')) NOT NULL DEFAULT 'active',
    refresh_token TEXT,
    refresh_token_exp TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS templates (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('active', 'inactive', 'deleted')) DEFAULT 'active',
    created_by BIGINT,
    type VARCHAR(20) CHECK (type IN ('online', 'offline')) DEFAULT 'online',
    level VARCHAR(20) CHECK (level IN ('beginner', 'intermediate', 'advanced')) DEFAULT 'beginner',
    language VARCHAR(20) DEFAULT 'English',
    description TEXT,
    image TEXT,
    price NUMERIC(10,2),
    discount NUMERIC(5,2),
    duration VARCHAR(50),
    capacity INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS courses (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    template_id BIGINT,
    schedule TEXT,
    status VARCHAR(20) DEFAULT 'active',
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS teacher_assignments (
    id BIGSERIAL PRIMARY KEY,
    id BIGINT,
    teacher_id BIGINT,
    course_id BIGINT,
    role VARCHAR(20) CHECK (role IN ('Teacher', 'Assistant')) NOT NULL
);

CREATE TABLE IF NOT EXISTS student_enrollments (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT,
    course_id BIGINT,
    status VARCHAR(20) CHECK (status IN ('enrolled', 'completed', 'dropped')) DEFAULT 'enrolled',
    progress INT DEFAULT 0,
    last_access TIMESTAMP,
    last_activity TIMESTAMP,
    last_activity_type VARCHAR(20) CHECK (last_activity_type IN ('quiz', 'assignment', 'forum', 'material')) DEFAULT 'quiz',
    enrolled_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS folders (
    id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    "order" INT DEFAULT 0,
    status VARCHAR(20) CHECK (status IN ('active', 'inactive', 'deleted')) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS materials (
    id BIGSERIAL PRIMARY KEY,
    teacher_id BIGINT,
    name VARCHAR(255),
    description TEXT,
    file_url TEXT,
    folder_id BIGINT,
    type VARCHAR(20) CHECK (type IN ('video', 'audio', 'document', 'link')) DEFAULT 'document',
    status VARCHAR(20) DEFAULT 'published',
    upload_date TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS forums CASCADE;
CREATE TABLE IF NOT EXISTS forums (
    id BIGSERIAL PRIMARY KEY,
    course_id BIGINT,
    title VARCHAR(255),
    description TEXT,
    status VARCHAR(20) CHECK (status IN ('open', 'closed')) DEFAULT 'open',
    created_by BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS threads (
    id BIGSERIAL PRIMARY KEY,
    forum_id BIGINT,
    title VARCHAR(255),
    content TEXT,
    status VARCHAR(20) CHECK (status IN ('open', 'closed')) DEFAULT 'open',
    views INT DEFAULT 0,
    replies INT DEFAULT 0,
    last_reply TIMESTAMP,
    created_by BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    thread_id BIGINT,
    author_id BIGINT,
    content TEXT,
    status VARCHAR(20) CHECK (status IN ('visible', 'hidden')) DEFAULT 'visible',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS feedbacks (
    id BIGSERIAL PRIMARY KEY,
    teacher_id BIGINT,
    student_id BIGINT,
    course_id BIGINT,
    comment TEXT,
    score FLOAT,
    status VARCHAR(20) CHECK (status IN ('pending', 'approved', 'rejected')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS registrations (
    id BIGSERIAL PRIMARY KEY,
    guest_name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(20),
    template_id BIGINT,
    course_id BIGINT,
    status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'cancelled')) DEFAULT 'pending',
    registration_date TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS payments (
    id BIGSERIAL PRIMARY KEY,
    registration_id BIGINT,
    amount NUMERIC(12,2),
    status VARCHAR(20) CHECK (status IN ('pending', 'completed', 'failed')) DEFAULT 'pending',
    payment_date TIMESTAMP DEFAULT NOW(),
    payment_method VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blogs (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255),
    content TEXT,
    author_id BIGINT,
    status VARCHAR(20) CHECK (status IN ('draft', 'published', 'archived')) DEFAULT 'draft',
    tags VARCHAR(255),
    image_url TEXT,
    comments_count INT DEFAULT 0,
    likes_count INT DEFAULT 0,
    views_count INT DEFAULT 0,
    comments_enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- ============================================
-- 2. ADD FOREIGN KEY CONSTRAINTS
-- ============================================
-- ===========================
-- FOREIGN KEY CONSTRAINTS
-- ===========================

-- Courses → Templates
ALTER TABLE courses
    ADD CONSTRAINT fk_courses_template
    FOREIGN KEY (template_id) REFERENCES templates(id);

-- Teacher Assignments → Courses, Users
ALTER TABLE teacher_assignments
    ADD CONSTRAINT fk_ta_course
    FOREIGN KEY (course_id) REFERENCES courses(id),
    ADD CONSTRAINT fk_ta_teacher
    FOREIGN KEY (teacher_id) REFERENCES users(id);

-- Student Enrollments → Courses, Users
ALTER TABLE student_enrollments
    ADD CONSTRAINT fk_enroll_student
    FOREIGN KEY (student_id) REFERENCES users(id),
    ADD CONSTRAINT fk_enroll_course
    FOREIGN KEY (course_id) REFERENCES courses(id);

-- Folders → Courses, Parent Folders
ALTER TABLE folders
    ADD CONSTRAINT fk_folder_course
    FOREIGN KEY (course_id) REFERENCES courses(id),
    ADD CONSTRAINT fk_folder_parent
    FOREIGN KEY (parent_id) REFERENCES folders(id);

-- Materials → Courses, Folders, Teachers
ALTER TABLE materials
    ADD CONSTRAINT fk_material_course
    FOREIGN KEY (course_id) REFERENCES courses(id),
    ADD CONSTRAINT fk_material_folder
    FOREIGN KEY (folder_id) REFERENCES folders(id),
    ADD CONSTRAINT fk_material_teacher
    FOREIGN KEY (teacher_id) REFERENCES users(id);

-- Forums → Courses
ALTER TABLE forums
    ADD CONSTRAINT fk_forum_course
    FOREIGN KEY (course_id) REFERENCES courses(id);

-- Threads → Forums, Users
ALTER TABLE threads
    ADD CONSTRAINT fk_thread_forum
    FOREIGN KEY (forum_id) REFERENCES forums(id),
    ADD CONSTRAINT fk_thread_creator
    FOREIGN KEY (created_by) REFERENCES users(id);

-- Posts → Threads, Users
ALTER TABLE posts
    ADD CONSTRAINT fk_post_thread
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    ADD CONSTRAINT fk_post_author
    FOREIGN KEY (author_id) REFERENCES users(id);

-- Feedbacks → Courses, Students, Teachers
ALTER TABLE feedbacks
    ADD CONSTRAINT fk_feedback_course
    FOREIGN KEY (course_id) REFERENCES courses(id),
    ADD CONSTRAINT fk_feedback_teacher
    FOREIGN KEY (teacher_id) REFERENCES users(id),
    ADD CONSTRAINT fk_feedback_student
    FOREIGN KEY (student_id) REFERENCES users(id);

-- Registrations → Courses, Templates
ALTER TABLE registrations
    ADD CONSTRAINT fk_registration_course
    FOREIGN KEY (course_id) REFERENCES courses(id),
    ADD CONSTRAINT fk_registration_template
    FOREIGN KEY (template_id) REFERENCES templates(id);

-- Payments → Registrations
ALTER TABLE payments
    ADD CONSTRAINT fk_payment_registration
    FOREIGN KEY (registration_id) REFERENCES registrations(id);

-- Blogs → Users (Admin)
ALTER TABLE blogs
    ADD CONSTRAINT fk_blog_author
    FOREIGN KEY (author_id) REFERENCES users(id);
