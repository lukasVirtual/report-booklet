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
      <h1 id="test"></h1>
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
import { io } from "socket.io-client";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "InstructorDefault",
  components: { BaseLayout },
  setup() {
    const who = ref<string[]>([]);
    const selected = ref("");
    const socket = io("http://localhost:7000");

    onMounted(async () => {
      const instructor = await loginService.getUser();
      const users = await dataService.GetAllUserForInstructor(instructor);
      who.value = users.map((v: { Name: string }) => v.Name);
      const elem = document.getElementById("test");
      socket.on("test", (message) => {
        console.log("recieved");
        if (elem) elem.innerHTML = message.date;
        console.log(message);
      });
    });

    return { who, selected };
  },
});
</script>
