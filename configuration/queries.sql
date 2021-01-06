   CREATE TABLE `countries` (
     `uuid` varchar(255) DEFAULT NULL,
     `name` varchar(255) DEFAULT NULL,
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

  CREATE TABLE `country_vaccines` (
     `uuid` varchar(255) DEFAULT NULL,
     `country_id` varchar(255) DEFAULT NULL,
     `vaccine_id` varchar(255) DEFAULT NULL,
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

    CREATE TABLE `users` (
    `uuid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `roll` bigint DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`uuid`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

    CREATE TABLE `user_docs` (
     `uuid` varchar(255) DEFAULT NULL,
     `user_id` varchar(255) DEFAULT NULL,
     `official_id` varchar(255) DEFAULT NULL,
     `official_id_type` varchar(255) DEFAULT NULL,
     `issuing_country_id` varchar(255) DEFAULT NULL,
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

    CREATE TABLE `user_vaccines` (
    `uuid` varchar(255) NOT NULL,
    `user_doc_id` varchar(255) NOT NULL,
    `vaccine_id` varchar(255) NOT NULL,
     `vaccine_authority_id` varchar(255) NOT NULL,
     `vaccination_date` varchar(255) NOT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`uuid`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

    CREATE TABLE `vaccines` (
     `uuid` varchar(255) DEFAULT NULL,
     `name` varchar(255) DEFAULT NULL,
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
