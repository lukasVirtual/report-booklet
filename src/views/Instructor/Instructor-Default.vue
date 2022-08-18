<template v-if="$route.name === 'Instructor'">
  <base-layout>
    <v-main style="margin: 1rem">
      <h1>Instructor</h1>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "InstructorDefault",
  components: { BaseLayout },
  setup() {
    return {};
  },
});
</script>

