<template v-if="$route.path === '/Berichtsheft'">
  <v-textarea
    @vnode-mounted="
      inputText = input;
      indexField = id;
      rowCounter = Number(rows);
      fields = rowCounter;
    "
    auto-grow
    spellcheck="true"
    :rows="rowCounter"
    @keyup.enter="rowCounter++"
    @keydown.delete="rowCounter > 1 ? rowCounter-- : (rowCounter = 1)"
    style="min-height: 60px"
    color="cyan"
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
import { defineComponent, ref, watchEffect } from "vue";
import { store } from "@/handler/store";
// import { timeStmp } from "./Text-Field.vue";

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
    if (store.input !== "") {
      console.log(store.input);
      emit("update");
    }
    // });
    const update = () => {
      if (
        store.input !== "" ||
        inputText.value == "" ||
        rowCounter.value !== fields.value
      ) {
        console.log(store.input);
        emit("update");
      }
    };

    const callback = () => {
      emit("update");
    };

    return { inputText, rowCounter, indexField, fields, update, callback };
  },
});
</script>
