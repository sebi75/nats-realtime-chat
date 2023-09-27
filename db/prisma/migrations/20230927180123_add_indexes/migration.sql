/*
  Warnings:

  - A unique constraint covering the columns `[username]` on the table `User` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE INDEX `idx_account_user_id` ON `Account`(`userId`);

-- CreateIndex
CREATE INDEX `idx_account_email` ON `Account`(`email`);

-- CreateIndex
CREATE UNIQUE INDEX `User_username_key` ON `User`(`username`);

-- CreateIndex
CREATE INDEX `idx_user_username` ON `User`(`username`);

-- CreateIndex
CREATE INDEX `idx_user_account_id` ON `User`(`account_id`);
