

CREATE DATABASE blackjack;

GRANT ALL PRIVILEGES ON blackjack.* TO 'serverConnect'@'%';

USE blackjack;

CREATE TABLE rooms (
  id         INT  UNSIGNED AUTO_INCREMENT,
  max_bet    DECIMAL(18,2) NOT NULL,
  min_bet    DECIMAL(18,2) NOT NULL,
  seats_num  INT           NOT NULL,
  last_round VARCHAR(64),
  PRIMARY KEY (id)
);

CREATE TABLE histories (
  id           INT UNSIGNED AUTO_INCREMENT,
  room_id      INT UNSIGNED NOT NULL,
  end_at       TIMESTAMP    NOT NULL,
  round_code   VARCHAR(64)  NOT NULL,
  dealer_cards JSON         NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (room_id) REFERENCES rooms (id)
);

CREATE TABLE seat_results (
  id           INT  UNSIGNED AUTO_INCREMENT,
  history_id   INT  UNSIGNED NOT NULL,
  seat_no      INT           NOT NULL,
  pile_0_bet   DECIMAL(18,2) NOT NULL,
  pile_0_pay   DECIMAL(18,2) NOT NULL,
  pile_0_cards JSON          NOT NULL,
  pile_1_bet   DECIMAL(18,2),
  pile_1_pay   DECIMAL(18,2),
  pile_1_cards JSON,
  insurance    DECIMAL(18,2),
  ins_pay      DECIMAL(18,2),
  PRIMARY KEY (id),
  FOREIGN KEY (history_id) REFERENCES histories (id)
);