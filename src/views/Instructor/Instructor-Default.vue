<template>
  <base-layout v-if="$route.name === 'Instructor'">
    <v-main style="margin: 1.5rem">
      <v-select
        :items="who"
        variant="underlined"
        style="max-width: 200px; max-height: 60px"
        v-model="selected"
      >
      </v-select>
      <v-container
        style="width: auto; height: auto"
        class="d-flex justify-center"
        fluid
      >
        <h1>Instructor</h1>
      </v-container>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted, ref } from "@vue/runtime-core";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "InstructorDefault",
  components: { BaseLayout },
  setup() {
    const who = ref<string[]>([]);
    const selected = ref("");

    onMounted(async () => {
      const instructor = await loginService.getUser();
      const users = await dataService.GetAllUserForInstructor(instructor);
      who.value = users.map((v: { Name: string }) => v.Name);
    });

    return { who, selected };
  },
});
</script>
