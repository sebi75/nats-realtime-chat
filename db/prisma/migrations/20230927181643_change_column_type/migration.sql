/*
  Warnings:

  - The `email_verified` column on the `Account` table would be dropped and recreated. This will lead to data loss if there is data in the column.

*/
-- AlterTable
ALTER TABLE `Account` DROP COLUMN `email_verified`,
    ADD COLUMN `email_verified` BOOLEAN NULL;
