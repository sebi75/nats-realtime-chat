/*
  Warnings:

  - You are about to drop the column `userId` on the `Account` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX `idx_account_user_id` ON `Account`;

-- AlterTable
ALTER TABLE `Account` DROP COLUMN `userId`;

-- CreateIndex
CREATE INDEX `idx_account_id` ON `Account`(`id`);
