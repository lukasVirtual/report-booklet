<template>
  <base-layout v-if="$route.path === '/Instructor'">
    <v-main style="margin: 1.5rem">
      <v-select
        :items="who"
        variant="underlined"
        style="max-width: 200px; max-height: 60px; min-width: 125px"
        v-model="selected"
      >
      </v-select>
      <v-card color="transparent" class="d-flex justify-center mb-6">
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn icon @click="switchPageLeft"
            ><v-icon>mdi-arrow-left</v-icon></v-btn
          >
          <v-btn icon @click="switchPageRight"
            ><v-icon>mdi-arrow-right</v-icon></v-btn
          >
        </v-card-actions>
      </v-card>

      <h1 class="d-flex justify-center">Instructor</h1>
      <v-container
        class="d-flex justify-center mb-6"
        v-for="date in uniqueDates"
        :key="`${who}.${date}.${month}.${year}`"
      >
        <viewing-field
          v-if="selected !== ''"
          :data="reports"
          :date="date"
        ></viewing-field>
      </v-container>
    </v-main>
    <template v-slot:navIcons>
      <v-btn icon id="menu-activator"
        ><v-icon style="width: 35px" size="22">mdi-email</v-icon></v-btn
      >
    </template>

    <v-menu activator="#menu-activator">
      <v-list max-height="300" min-width="300">
        <v-list-item v-for="(item, index) in logs" :key="index" :value="index">
          <template v-slot:prepend>
            <v-icon icon="mdi-bell" size="20" color="error"></v-icon>
          </template>
          <v-list-item-title style="font-size: 13px; color: dimgray">{{
            item.toUpperCase()
          }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </base-layout>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted, ref, watch } from "@vue/runtime-core";
import { io } from "socket.io-client";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import ViewingField from "./Viewing-Field.vue";

export default defineComponent({
  name: "InstructorDefault",
  components: { BaseLayout, ViewingField },
  setup() {
    const who = ref<string[]>([]);
    const selected = ref("");
    const socket = io("http://localhost:7000");
    const reports = ref<any[]>([]);
    const uniqueDates = ref<string[]>([]);
    const date = new Date();
    const month = ref(date.getMonth() + 1);
    const year = ref(date.getFullYear());
    const logs = ref<string[]>([]);

    onMounted(async () => {
      const instructor = await loginService.getUser();
      const users = await dataService.GetAllUserForInstructor(instructor);
      const retrievedLogs: string[] = await dataService.RetrieveLog(instructor);
      logs.value = retrievedLogs;
      console.warn(logs);
      who.value = users.map((v: { Name: string }) => v.Name);

      // console.log(currentDate.value);
      socket.on("test", async (user) => {
        const retrievedLogs = await dataService.RetrieveLog(instructor);
        logs.value = retrievedLogs;
        //   console.debug(`Data Uploaded from ${user}`);
        // if (elem) elem.innerHTML = object[0].Date;
        // console.log(object);
        // TODO Write Status/Notification in Backlog
        // await dataService.RetrieveSubmittedData(name);
      });
    });

    const switchPageLeft = () => {
      if (month.value > 1) month.value--;
    };

    const switchPageRight = () => {
      if (month.value < 12) month.value++;
    };

    watch([selected, month], async ([newSelection, newMonth]) => {
      try {
        uniqueDates.value = [];
        reports.value = await dataService.RetrieveSubmittedData(
          newSelection,
          newMonth.toString()
        );
        const cacheDates = [];
        for (const date of reports.value) {
          cacheDates.unshift(date.Date);
        }
        uniqueDates.value = [...new Set(cacheDates)].sort();
      } catch (e) {
        reports.value = [];
        uniqueDates.value = [`1.${month.value}.${year.value}`];
      }
    });

    return {
      who,
      selected,
      reports,
      month,
      uniqueDates,
      switchPageLeft,
      switchPageRight,
      date,
      logs,
      year,
    };
  },
});
</script>
