import { reactive } from "vue";

type Store = {
  id: undefined | number;
  date: string;
  input: string;
  time: string;
  rows: number;
};

export const store: Store = {
  id: undefined,
  date: "",
  input: "",
  time: "",
  rows: 1,
};

export const storage: Store[] = [];
