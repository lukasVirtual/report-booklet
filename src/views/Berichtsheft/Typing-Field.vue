<template v-if="$route.path === '/Berichtsheft'">
  <v-text-field :auto-grow="true" spellcheck="true"></v-text-field>
  <timer-text-field></timer-text-field>
</template>

<script lang="ts">
import TimerTextField from "./Timer-TextField.vue";
import { defineComponent } from "vue";

export default defineComponent({
  name: "TypingField",
  components: { TimerTextField },

  setup() {
    return {};
  },
});
</script>