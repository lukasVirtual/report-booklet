<template v-if="$route.path === '/Berichtsheft'">
  <v-textarea
    @vnode-mounted="
      inputText = input;
      indexField = id;
    "
    auto-grow
    spellcheck="true"
    :rows="rowCounter"
    @keyup.enter="rowCounter++"
    @keydown.delete="rowCounter > 1 ? rowCounter-- : (rowCounter = 1)"
    style="min-height: 60px"
    color="cyan"
    v-model="inputText"
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
import { defineComponent, onMounted, ref, watchEffect } from "vue";
import { store } from "@/handler/store";
// import { timeStmp } from "./Text-Field.vue";

export default defineComponent({
  name: "TypingField",
  components: { TimerTextField },
  props: {
    input: String,
    time: String,
    id: Number,
  },
  emits: ["update"],

  setup(props, { emit }) {
    const inputText = ref<string | undefined>("");
    const indexField = ref<number | undefined>(0);
    const rowCounter = ref(1);

    onMounted(() => {
      //   console.log(indexField.value);
      console.log(inputText.value);
    });

    watchEffect(() => {
      store[0].input = inputText.value as string;
      store[0].id = indexField.value;
      // console.log(inputText.value);
      if (store[0].input !== "" && store[0].time !== "") {
        emit("update");
      }
    });

    const callback = () => {
      emit("update");
    };

    return { inputText, rowCounter, indexField, callback };
  },
});
</script>
