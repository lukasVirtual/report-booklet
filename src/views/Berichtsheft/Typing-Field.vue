<template>
  <v-textarea
    v-if="$route.path === '/Berichtsheft'"
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
import { store } from "@/handler/store";

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

    watchEffect(() => {
      store.input = inputText.value as string;
      store.id = indexField.value;
      store.rows = rowCounter.value as number;
      store.time = "00:00";
      console.log(fields.value, rowCounter.value);
    });

    const update = () => {
      if (
        store.input !== "" ||
        inputText.value === "" ||
        rowCounter.value !== fields.value
      ) {
        console.debug(
          "testing typing field: ",
          store.input,
          inputText.value,
          rowCounter.value,
          fields.value
        );
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
