<template>
  <v-textarea
    v-if="$route.path === '/Berichtsheft' || $route.path === '/Instructor'"
    @vnode-mounted="
      inputText = input;
      indexField = id;
      fields = rowCounter;
    "
    rows="1"
    auto-grow
    spellcheck
    color="cyan"
    filled
    v-model="inputText"
    @blur="update"
  ></v-textarea>
  <div>
    <timer-text-field
      @update-time="callback"
      :timeStamp="time"
      :index="id"
    ></timer-text-field>
  </div>
</template>

<script lang="ts">
import TimerTextField from "./Timer-TextField.vue";
import { defineComponent, ref, watchEffect } from "@vue/runtime-core";
import { useStore } from "@/handler/store";
import { storeToRefs } from "pinia";

export default defineComponent({
  name: "TypingField",
  components: { TimerTextField },
  props: {
    input: String,
    time: String,
    id: Number,
    rows: Number,
  },
  emits: ["update"],

  setup(props, { emit }) {
    const inputText = ref<string | undefined>("");
    const indexField = ref<number | undefined>(0);
    const rowCounter = ref<number>(1);
    const fields = ref(rowCounter.value);
    const initStore = useStore();
    const { store } = storeToRefs(initStore);

    watchEffect(() => {
      store.value.input = inputText.value as string;
      store.value.id = indexField.value;
      store.value.rows = rowCounter.value as number;
      store.value.time = "00:00";
    });

    const update = () => {
      if (
        store.value.input !== "" ||
        inputText.value === "" ||
        rowCounter.value !== fields.value
      ) {
        emit("update");
      }
    };

    const callback = () => {
      emit("update");
    };

    return {
      inputText,
      rowCounter,
      indexField,
      fields,
      update,
      callback,
    };
  },
});
</script>
