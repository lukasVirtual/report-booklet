import { defineStore } from "pinia";
import { ref } from "vue";

type Store = {
  id: undefined | number;
  date: string;
  input: string;
  time: string;
  rows: number;
};

export const useStore = defineStore("text-field", () => {
  const store = ref<Store>({
    id: undefined,
    date: "",
    input: "",
    time: "",
    rows: 1,
  });
  return { store };
});
