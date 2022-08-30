<template>
  <div v-if="$route.path == '/Berichtsheft'" class="text-center">
    <v-menu
      anchor="bottom"
      transition="slide-x-transition"
      @vnode-mounted="
        date = dateTextField;
        status = statusTextField;
      "
    >
      <template v-slot:activator="{ props }">
        <v-btn
          @click="test"
          v-bind="props"
          icon
          style="margin-bottom: 30px"
          elevation="20"
          ><v-icon size="30">mdi-text-box-outline</v-icon>
        </v-btn>
        <v-btn icon size="15">{{ stateTrue?.length ?? 0 }}</v-btn>
      </template>
      <v-list width="400px" height="500">
        <v-list-item v-for="(quali, idx) in qualis" :key="idx">
          <v-list-item-title>
            <v-list-item-title>{{ quali.Text }}</v-list-item-title>
          </v-list-item-title>
          <v-spacer></v-spacer>
          <v-list-item-action>
            <v-checkbox
              v-model="quali.State"
              @click="changeState(idx)"
              color="deep-purple accent-4"
            ></v-checkbox>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  name: "QualificationsDefault",

  props: {
    dateTextField: String,
    statusTextField: String,
  },
  setup() {
    const date = ref<string | undefined>("");
    const qualis = ref();
    const status = ref<string | undefined>("");
    const stateTrue = ref();

    const qualifications = [
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
      {
        text: "the first qualification",
        state: false,
      },
    ];

    onMounted(async () => {
      qualis.value = await dataService.getQualifications(date.value as string);
      stateTrue.value = qualis.value?.filter((v: any) => v.State == true);
    });

    const test = async () => {
      console.log(date.value);
      console.log(status.value);
      // console.error(qualis.value);
      if (qualis.value?.length < 16 || qualis.value == undefined) {
        await dataService.insertQualifications(
          qualifications,
          date.value as string
        );
        console.log("empty");
      }
      qualis.value = await dataService.getQualifications(date.value as string);
      // console.warn(qualis.value);
    };

    const changeState = async (idx: number) => {
      // qualifications[idx].state = !qualifications[idx].state;
      qualis.value[idx].State = !qualis.value[idx].State;
      await dataService.insertQualifications(
        qualis.value,
        date.value as string
      );
      qualis.value = await dataService.getQualifications(date.value as string);
      stateTrue.value = qualis.value?.filter((v: any) => v.State == true);
    };

    return {
      test,
      changeState,
      status,
      date,
      stateTrue,
      qualifications,
      qualis,
    };
  },
});
</script>
