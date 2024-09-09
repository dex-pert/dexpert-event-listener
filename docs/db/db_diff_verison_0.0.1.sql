CREATE TABLE `user_launch_tx`  (
                                   `Id` int NOT NULL AUTO_INCREMENT,
                                   `uid` bigint NOT NULL COMMENT '用户id',
                                   `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '合约地址',
                                   `chain_id` int NOT NULL COMMENT '链Id',
                                   `pair_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '对币地址',
                                   `fee` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易手续费',
                                   `fee_token_symbol` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易手续费token的symbol',
                                   `timestamp` timestamp NOT NULL COMMENT '交易时间',
                                   `type_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类别: luanch',
                                   `chain_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                   PRIMARY KEY (`Id`) USING BTREE,
                                   UNIQUE INDEX `uix_chain_contract`(`chain_id` ASC, `contract_address` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 77 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

CREATE TABLE `user_swap_tx`  (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `uid` bigint NOT NULL COMMENT '用户编号',
                                 `address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '引用user_wallet表的address',
                                 `tx` varchar(66) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易哈希值',
                                 `token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '第一个代币的名称',
                                 `amount_in` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '第一个代币的数量',
                                 `token_in` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '第一个代币的地址',
                                 `transaction_time` timestamp NOT NULL COMMENT '时间',
                                 `fee` varchar(66) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手续费',
                                 `block_number` int NOT NULL COMMENT '区块Id',
                                 `chain_id` int NOT NULL COMMENT '链id',
                                 `chain_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链名',
                                 `swap_type` int NOT NULL DEFAULT 0 COMMENT '交易类型0：swap，1：Sniper',
                                 `token_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '代币名称',
                                 `decimal` int NOT NULL DEFAULT 18,
                                 `fee_token_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '费率token',
                                 `fee_decimal` int NOT NULL,
                                 `fee_token_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 INDEX `fk_user_wallet_address`(`address` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 261 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录用户swap产生的交易' ROW_FORMAT = Dynamic;

CREATE TABLE `user_transaction`  (
                                     `id` int NOT NULL AUTO_INCREMENT COMMENT '记录编号',
                                     `uid` bigint NOT NULL COMMENT '用户id',
                                     `tid` int NOT NULL COMMENT '关联的table的 id',
                                     `swap_type` tinyint NOT NULL COMMENT '交易类型 0:swap 1:sniper 99: launch',
                                     `volume` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.00' COMMENT '交易量',
                                     `timestamp` timestamp NOT NULL COMMENT '交易时间',
                                     `swap_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '交易类型的名字',
                                     `chain_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链名',
                                     `chain_id` int NOT NULL COMMENT '链id',
                                     `fee` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手续费',
                                     `fee_token_symbol` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手续费symbol',
                                     `fee_token_decimal` int NOT NULL COMMENT '手续费token的decimal',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 202 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;