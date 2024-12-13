-- Insert data into the 'users' table
INSERT INTO users (username, email) VALUES
    ('john_doe', 'john.doe@example.com'),
    ('jane_smith', 'jane.smith@example.com'),
    ('alice_wonder', 'alice.wonder@example.com');

-- Insert data into the 'meetups' table
INSERT INTO meetups (name, description, user_id) VALUES
    ('Tech Enthusiasts Meetup', 'A meetup for people passionate about technology.', 1),
    ('Startup Founders Forum', 'A space for startup founders to network and share ideas.', 2),
    ('Book Lovers Club', 'Monthly meetup for book discussions and reviews.', 3);