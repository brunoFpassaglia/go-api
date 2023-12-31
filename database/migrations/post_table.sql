USE SOCIALNETWORK ;
DROP TABLE IF EXISTS POST ;

CREATE TABLE POSTS(
    ID INT AUTO_INCREMENT PRIMARY KEY,
    TITLE VARCHAR(50) not null,
    CONTENT VARCHAR(80) NOT NULL  ,
    AUTHOR_ID INT NOT NULL  ,
    FOREIGN KEY (AUTHOR_ID) REFERENCES USERS(ID) ON DELETE CASCADE,
    LIKES INT NOT NULL DEFAULT 0  ,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;
