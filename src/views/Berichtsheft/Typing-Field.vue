<template v-if="$route.path === '/Berichtsheft'">
  <v-text-field
    :auto-grow="true"
    spellcheck="true"
    v-model="inputText"
  ></v-text-field>
  <timer-text-field></timer-text-field>
</template>

<script lang="ts">
import TimerTextField from "./Timer-TextField.vue";
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "TypingField",
  components: { TimerTextField },

  setup() {
    const inputText = ref("");

    return { inputText };
  },
});
</script>
