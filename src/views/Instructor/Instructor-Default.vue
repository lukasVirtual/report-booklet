<template>
  <base-layout v-if="$route.name === 'Instructor'">
    <v-main style="margin: 1.5rem">
      <v-select
        :items="who"
        variant="underlined"
        style="max-width: 200px; max-height: 60px; min-width: 125px"
        v-model="selected"
      >
      </v-select>
      <!-- <v-card color="transparent" class="d-flex justify-center mb-6">
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn icon @click="switchPageLeft"
            ><v-icon>mdi-arrow-left</v-icon></v-btn
          >
          <v-btn @click="today" style="font-size: 20px">Heute</v-btn>
          <v-btn icon @click="switchPageRight"
            ><v-icon>mdi-arrow-right</v-icon></v-btn
          >

          <v-btn :disabled="toggleStage" @click="submit" color="blue"
            >Submit</v-btn
          >

          <h5 style="color: grey">{{ time }}</h5>
        </v-card-actions>
      </v-card> -->
      <h1>Instructor</h1>
      <v-container fluid>
        <v-card
          style="
            width: 1100px;
            height: 500px;
            border: 1px solid black;
            overflow-x: hidden;
            border-radius: 10px;
          "
          max-height="500"
          class="mt-4 overflow-y-auto"
        >
          <v-row dense>
            <div style="width: 5px"></div>
            <v-row style="max-height: 30px">
              <v-col cols="12" md="12">
                <div v-for="report in reports" :key="report.Input">
                  <p>{{ report.Date }}</p>
                  <v-card-actions>
                    <typing-field
                      :input="report.Input"
                      :time="report.Time"
                    ></typing-field>
                  </v-card-actions>
                </div>
              </v-col>
              <v-spacer></v-spacer>
            </v-row>
          </v-row>
        </v-card>
      </v-container>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted, ref, watch } from "@vue/runtime-core";
import { io } from "socket.io-client";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import TypingField from "../Berichtsheft/Typing-Field.vue";
// import uniqBy from "lodash/uniqBy";

export default defineComponent({
  name: "InstructorDefault",
  components: { BaseLayout, TypingField },
  setup() {
    const who = ref<string[]>([]);
    const selected = ref("");
    const socket = io("http://localhost:7000");
    const reports = ref<any>([]);

    onMounted(async () => {
      const instructor = await loginService.getUser();
      const users = await dataService.GetAllUserForInstructor(instructor);
      who.value = users.map((v: { Name: string }) => v.Name);
      if (selected.value !== "") {
        try {
          reports.value = await dataService.RetrieveSubmittedData(
            selected.value
          );
        } catch (e) {
          console.warn(e);
        }
      }
      // const elem = document.getElementById("test");
      socket.on("test", (user) => {
        console.debug(`Data Uploaded from ${user}`);
        // if (elem) elem.innerHTML = object[0].Date;
        // console.log(object);

        // TODO Write Status/Notification in Backlog

        // await dataService.RetrieveSubmittedData(name);
      });
    });

    watch(selected, async (newSelection: string) => {
      try {
        reports.value = await dataService.RetrieveSubmittedData(newSelection);
        // const uniqueDates = computed(() => {
        //   return uniqBy(reports, (e: any) => {
        //     return e.Date;
        //   });
        // });
        console.log(reports.value);
        // console.log(uniqueDates.value);
      } catch (e) {
        console.error(e);
      }
    });

    return { who, selected, reports };
  },
});
</script>
