ALTER TABLE  post add
    CONSTRAINT fk_post_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(id);

ALTER TABLE  comment add
    CONSTRAINT fk_comment_user_id
        FOREIGN KEY (user_id)
            REFERENCES users(id);

ALTER TABLE  comment add
    CONSTRAINT fk_comment_post_id
        FOREIGN KEY (post_id)
            REFERENCES post(id);