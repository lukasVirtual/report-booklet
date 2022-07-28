<template v-if="$route.path === '/Berichtsheft'">
  <v-textarea
    @vnode-mounted="inputText = input"
    auto-grow
    spellcheck="true"
    :rows="rowCounter"
    @keyup.enter="rowCounter++"
    @keydown.delete="rowCounter > 1 ? rowCounter-- : (rowCounter = 1)"
    style="min-height: 60px"
    color="cyan"
    v-model="inputText"
  ></v-textarea>
  <timer-text-field></timer-text-field>
</template>

<script lang="ts">
import TimerTextField from "./Timer-TextField.vue";
import { defineComponent, ref, watchEffect } from "vue";
import { store } from "@/Auth/store";

export default defineComponent({
  name: "TypingField",
  components: { TimerTextField },
  props: {
    input: String,
  },

  setup() {
    const inputText = ref<string | undefined>("");
    const rowCounter = ref(1);

    watchEffect(() => {
      store[0].input = inputText.value as string;
    });

    return { inputText, rowCounter };
  },
});
</script>
