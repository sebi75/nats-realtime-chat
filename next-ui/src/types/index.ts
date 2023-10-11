export type Json = {
  [key: string]: string | number | boolean | Json | JsonArray;
};

export type JsonArray = Array<Json>;

export enum HttpMethod {
  Post = "POST",
  Get = "GET",
  Put = "PUT",
  Delete = "DELETE",
}
