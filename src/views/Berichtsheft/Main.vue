<template>
  <base-layout v-if="$route.path === '/Berichtsheft'">
    <v-main>
      <v-card
        :disabled="toggleStage"
        color="transparent"
        class="d-flex justify-center mb-6"
      >
        <v-card-actions>
          <v-btn icon @click="switchPageLeft"
            ><v-icon>mdi-arrow-left</v-icon></v-btn
          >
          <v-btn @click="today" style="font-size: 20px">Heute</v-btn>
          <v-btn icon @click="switchPageRight"
            ><v-icon>mdi-arrow-right</v-icon></v-btn
          >
          <v-btn @click="submit" color="blue">Submit</v-btn>
          <h6>{{ time }}</h6>
        </v-card-actions>
      </v-card>
      <v-container
        class="d-flex justify-center mb-6"
        v-for="i in daysOfMonth"
        :key="i"
        ref="container"
        id="test"
      >
        <!-- <text-field :propsDate="`${i}.${currMonth}.${currYear}`"></text-field> -->
      </v-container>
    </v-main>

    <router-view />
  </base-layout>
</template>

<script lang="ts">
import {
  createApp,
  defineComponent,
  getCurrentInstance,
  onMounted,
  ref,
} from "vue";
import TextField from "./Text-Field.vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  components: { BaseLayout },
  name: "MainPage",

  setup() {
    const date = new Date();

    const { appContext } = getCurrentInstance() as any;
    console.log(appContext);
    let daysOfMonth = ref(31);
    let toggleStage = ref(false);
    let currMonth = ref(date.getMonth() + 1);
    const currYear = ref(date.getFullYear());
    let time = ref<string | null>(null);

    onMounted(() => {
      createComponentInstance();
    });
    const computeDays = () => {
      if (
        currMonth.value === 4 ||
        currMonth.value === 6 ||
        currMonth.value === 9 ||
        currMonth.value === 11
      ) {
        daysOfMonth.value--;
      } else if (currMonth.value !== 2) {
        daysOfMonth.value = 31;
      }

      if (currMonth.value == 2) {
        daysOfMonth.value = daysOfMonth.value - 3;

        if (currYear.value % 4 == 0) daysOfMonth.value++;
        if (currYear.value % 100 == 0) daysOfMonth.value--;
        if (currYear.value % 400 == 0) daysOfMonth.value++;
      }
    };
    let props: { propsDate: string } = { propsDate: "" };
    const createComponentInstance = async () => {
      // for (let i = 1; i <= daysOfMonth.value; i++) {
      props = {
        propsDate: `${daysOfMonth.value}.${currMonth.value}.${currYear.value}`,
      };
      const app = createApp((await import("./Text-Field.vue")).default, props);
      // localStorage.setItem(props.propsDate, props.propsDate);
      // console.log(app._context.app._uid++);
      Object.assign(app._context, appContext);
      app.mount("#test");
      // }
    };
    const switchPageRight = () => {
      currMonth.value++;
      if (currMonth.value > 12) {
        currMonth.value = 1;
        currYear.value++;
      }
      computeDays();
      const obj = localStorage.getItem(props.propsDate);

      createComponentInstance();
    };
    const switchPageLeft = async () => {
      if (currYear.value >= date.getFullYear()) {
        currMonth.value--;
        if (currMonth.value < 1) {
          currMonth.value = 12;
          currYear.value--;
        }
      }
      computeDays();
      createComponentInstance();
    };

    const today = () => {
      currMonth.value = date.getMonth() + 1;
      currYear.value = date.getFullYear();
      createComponentInstance();
    };

    const submit = () => {
      toggleStage.value == false
        ? (toggleStage.value = true)
        : (toggleStage.value = false);

      time.value = `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;

      localStorage.setItem("stage", JSON.stringify(document.body.innerHTML));
    };

    return {
      daysOfMonth,
      currMonth,
      currYear,
      toggleStage,
      time,
      submit,
      switchPageRight,
      switchPageLeft,
      today,
    };
  },
});
</script>