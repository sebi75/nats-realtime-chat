export type Query<T> = {
  data: T;
  meta: {
    count: number;
  };
};
