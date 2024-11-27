
CREATE TABLE song (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      song_name VARCHAR(100) NOT NULL,
                      group_name VARCHAR(100) NOT NULL
);

INSERT INTO song (song_name, group_name) VALUES
                                             ('Imagine', 'John Lennon'),
                                             ('Hey Jude', 'The Beatles'),
                                             ('Billie Jean', 'Michael Jackson');

