package domain

/**
model User {
    id        String  @id @default(uuid())
    accountId String  @unique @map("account_id")
    username  String  @map("username")
    imageUrl  String? @map("image_url")

    account  Account      @relation(fields: [accountId], references: [id], onDelete: Cascade)
    chats    ChatMember[]
    messages Message[]
}
*/

type User struct {
	Id        string `json:"id"`
	AccountId string `json:"account_id"`
	Username  string `json:"username"`
	ImageUrl  string `json:"image_url"`
}
