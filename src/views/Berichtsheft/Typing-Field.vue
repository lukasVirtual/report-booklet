<template v-if="$route.path === '/Berichtsheft'">
  <v-textarea
    auto-grow
    spellcheck="true"
    :rows="rowCounter"
    @keyup.enter="rowCounter++"
    @keydown.delete="rowCounter > 1 ? rowCounter-- : (rowCounter = 1)"
    style="min-height: 60px"
    v-model="inputText"
  ></v-textarea>
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
    const rowCounter = ref(1);

    return { inputText, rowCounter };
  },
});
</script>
