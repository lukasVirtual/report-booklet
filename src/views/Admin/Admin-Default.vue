<template>
  <base-layout v-if="$route.name === 'Admin'">
    <v-main style="margin: auto">
      <h2 class="font-weight-bold text-center mt-5">Admin Center</h2>

      <v-dialog v-model="dialog" width="600px">
        <template v-slot:activator="{ attrs }">
          <v-btn icon @click="dialog = true" v-bind="attrs"><v-icon>mdi-plus</v-icon></v-btn>
        </template>
        <v-card width="550" height="450" rounded>
          <v-card-title class="text-h5 grey lighten-2">
            Register new User
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-row>
              <v-col>
                <v-text-field label="Name" variant="underlined" color="cyan" v-model="input.name"
                  :rules="[rules.required]"></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-text-field label="Password" variant="underlined" color="cyan" type="password"
                  v-model="input.password" :rules="[rules.required, rules.minLen]"></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-select :items="roleItems" label="Role" variant="underlined" color="cyan" v-model="input.role"
                  :rules="[rules.required]"></v-select>
              </v-col>
            </v-row>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-btn color="red" text @click="dialog = false"> Cancel </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="blue" text @click="addNewUser"> Register </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog v-model="relationDialog" width="600px">
        <template v-slot:activator="{ attrs }">
          <v-btn style="margin-left: 10px" icon v-bind="attrs" @click="
  relationDialog = true;
loadData();
          "><v-icon>mdi-relation-one-or-many-to-only-one</v-icon></v-btn>
        </template>

        <v-card width="550" height="450" rounded>
          <v-card-title class="text-h5 grey lighten-2">
            Assign User to Instructor
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-row>
              <v-col>
                <v-select :items="instructors" label="Select Instructor" variant="underlined" color="cyan"
                  v-model="chef"></v-select>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-select :items="users" label="Select User" variant="underlined" color="cyan" item-text="role"
                  item-value="name" v-model="selected" multiple chips>
                </v-select>
              </v-col>
            </v-row>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-btn color="red" text @click="
  relationDialog = false;
chef = '';
selected = [];
            ">
              Cancel
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="blue" text @click="assignUser"> Assign </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-table style="width: 1100px; max-width: 1400px; border-radius: 10px" class="mt-5">
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Role</th>
            <th class="text-left"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(elem, idx) in resultArr" :key="idx">
            <td>{{ elem.name }}</td>
            <td>{{ elem.role }}</td>
            <td class="text-right">
              <v-btn style="margin: 5px" v-if="elem.role === 'instructor'" @click="
  showRelations(elem.name);
toggled = !toggled;
instructor = elem.name;
              "><v-icon size="20">mdi-cogs</v-icon></v-btn>

              <v-btn v-if="elem.role !== 'admin'" @click="deleteElemAt(elem, idx)"><v-icon size="20"
                  color="red">mdi-trash-can</v-icon></v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>

      <div style="height: 50px"></div>
      <v-select v-if="toggled" style="max-width: 300px" multiple chips density="compact" closable-chips
        v-model="belongingUsers" :items="belongingUsers">
      </v-select>
      <!-- <v-card width="550" height="450" rounded v-if="toggled">
        <v-select
          multiple
          chips
          closable-chips
          v-model="belongingUsers"
          :items="belongingUsers"
        ></v-select>
      </v-card> -->
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import type { DataServiceInterface } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler"
// import type { LoginServiceInterface } from "@/handler/loginHandler";

import {
  defineComponent,
  onMounted,
  reactive,
  ref,
  inject,
} from "@vue/runtime-core";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import { useToast } from "vue-toastification";

export default defineComponent({
  name: "AdminDefault",
  components: { BaseLayout },

  setup() {
    // const loginService = inject(
    //   "provide-login-service"
    // ) as LoginServiceInterface;
    const dataService = inject("provide-data-service") as DataServiceInterface;

    const role = ref("");
    const toast = useToast();
    let instructors = ref<string[]>([]);
    const instructor = ref("");
    let users = ref<string[]>([]);
    const selected = ref([]);
    const chef = ref("");
    const belongingUsers = ref<any>([]);
    const previousBelongingUsers = ref<any>([]);
    const toggled = ref(false);
    const data = ref<any>([]);

    const resultArr = ref<[{ name: string; role: string }]>([
      { name: "", role: "" },
    ]);

    const rules = {
      required: (value: string) => !!value || "Required",
      minLen: (v: string) => v.length >= 8 || "Min. 8 Characters",
    };

    const input = reactive({
      name: "",
      password: "",
      role: "",
    });

    const dialog = ref(false);
    const relationDialog = ref(false);
    const counter = ref(1);
    const roleItems = ["admin", "instructor", "user"];

    const loadData = async (): Promise<void> => {
      data.value = await dataService.getAllData();
      instructors.value = [];
      users.value = [];
      let getInstructors = data.value.filter(
        (v: { Role: string }) => v.Role === "instructor" || v.Role === "user"
      );
      for (const p of getInstructors) {
        if (p.Role === "instructor" && !instructors.value.includes(p.name))
          instructors.value.push(p.Name);
        else if (
          p.Role === "user" &&
          p.Parent_id === 0 &&
          !users.value.includes(p.name)
        )
          users.value.push(p.Name);
      }
    };

    onMounted(async () => {
      role.value = (await loginService.getUserRole()) as unknown as string;
      data.value = await dataService.getAllData();
      for (const elem of data.value) {
        resultArr.value.push({
          name: elem.Name,
          role: elem.Role,
        });
      }

      resultArr.value.map((value, idx) => {
        if (value.name == "") resultArr.value.splice(idx, 1);
      });
    });

    const addNewUser = async () => {
      if (
        input.name !== "" &&
        input.password.length >= 8 &&
        input.role !== ""
      ) {
        loginService.register(input.name, input.password, input.role);
        resultArr.value.push({ name: input.name, role: input.role });
        toast.success("successfully registered new user");
        input.name = "";
        input.password = "";
        input.role = "";
      } else toast.error("Some fields are Empty or incorrect entered");
      dialog.value = false;
    };

    const deleteElemAt = async (elem: any, idx: number) => {
      if (
        !confirm(
          "You sure you want to delete? Deleted Items will be gone forever"
        )
      )
        return;
      try {
        if (elem.role === "instructor") {
          await dataService.DelteUsersFromInstructor(elem.name);
        }
        await dataService.deleteItem(elem.name);
        resultArr.value.splice(idx, 1);

        toast.success("successfully deleted user");
      } catch (e) {
        console.log(e);
      }
      await dataService.getAllData();
    };

    const assignUser = async () => {
      if (selected.value?.length <= 0 || chef.value === "") {
        toast.error("Some fields are empty please choose a value");
        relationDialog.value = false;
        selected.value = [];
        chef.value = "";
        return;
      }

      for (const user of selected.value) {
        await dataService.AssignUserToInstructor(chef.value, user);
      }
      toast.success("Successfully assigned User/Users to Instructor");
      relationDialog.value = false;

      selected.value = [];
      chef.value = "";
    };

    const showRelations = async (chef: string) => {
      belongingUsers.value = [];
      const out: any[] = await dataService.GetAllUserForInstructor(chef);
      for (const elem of out) {
        belongingUsers.value.push(elem.Name);
      }

      previousBelongingUsers.value = [...belongingUsers.value];
    };

    //TODO delete single user when clicked

    // watchEffect(() => {
    //   let diff = previousBelongingUsers.value.filter(
    //     (v: any) => !belongingUsers.value.includes(v)
    //   );
    //   console.log(instructor.value);
    //   for (const user of diff) {
    //     dataService.DelteSingleUserFromInstructor(user);
    //   }
    // });
    return {
      resultArr,
      dialog,
      roleItems,
      input,
      rules,
      role,
      relationDialog,
      instructors,
      users,
      counter,
      selected,
      chef,
      belongingUsers,
      toggled,
      instructor,
      showRelations,
      assignUser,
      addNewUser,
      deleteElemAt,
      loadData,
    };
  },
});
</script>
