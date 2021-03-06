-- MySQL Script generated by MySQL Workbench
-- Sat 16 Sep 2017 08:41:15 PM EDT
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema ulbora_api_gateway
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema ulbora_api_gateway
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `ulbora_api_gateway` DEFAULT CHARACTER SET utf8 ;
USE `ulbora_api_gateway` ;

-- -----------------------------------------------------
-- Table `ulbora_api_gateway`.`client`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ulbora_api_gateway`.`client` (
  `client_id` BIGINT NOT NULL,
  `api_key` VARCHAR(500) NOT NULL,
  `enabled` TINYINT(1) NOT NULL,
  `level` VARCHAR(45) NOT NULL DEFAULT 'small',
  PRIMARY KEY (`client_id`),
  UNIQUE INDEX `client_id_UNIQUE` (`client_id` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ulbora_api_gateway`.`rest_route`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ulbora_api_gateway`.`rest_route` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `route` VARCHAR(100) NOT NULL,
  `client_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`, `client_id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_rest_route_client1_idx` (`client_id` ASC),
  UNIQUE INDEX `index4` (`route` ASC, `client_id` ASC),
  CONSTRAINT `fk_rest_route_client1`
    FOREIGN KEY (`client_id`)
    REFERENCES `ulbora_api_gateway`.`client` (`client_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ulbora_api_gateway`.`route_url`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ulbora_api_gateway`.`route_url` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL DEFAULT 'blue',
  `url` VARCHAR(200) NOT NULL,
  `active` TINYINT(1) NOT NULL,
  `rest_route_id` BIGINT NOT NULL,
  `rest_route_client_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`, `rest_route_id`, `rest_route_client_id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_route_url_rest_route1_idx` (`rest_route_id` ASC, `rest_route_client_id` ASC),
  CONSTRAINT `fk_route_url_rest_route1`
    FOREIGN KEY (`rest_route_id` , `rest_route_client_id`)
    REFERENCES `ulbora_api_gateway`.`rest_route` (`id` , `client_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
