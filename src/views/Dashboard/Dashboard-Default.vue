<template>
  <base-layout v-if="$route.path === '/Dashboard'">
    <v-main style="margin: 20px">
      <h3>Dashboard</h3>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "DashboardDefault",
  components: { BaseLayout },
  setup() {
    return {};
  },
});
</script>