<template>
  <v-app>
    <v-main
      v-if="$route.path == '/Login'"
      class="text-center"
      style="display: flex; align-items: center; margin: auto"
    >
      <v-form
        tile
        ref="valid"
        style="
          border: 1px solid black;
          height: 500px;
          width: 800px;
          display: flex;
          justify-content: center;
          align-items: center;
        "
      >
        <v-row justify="center">
          <v-card-title>Login</v-card-title>
          <v-col cols="12">
            <v-text-field
              label="Username"
              type="text"
              style="padding-left: 25%; padding-right: 25%"
              :rules="[rules.required]"
              variant="underlined"
              v-model="validation.username"
              single-line
            ></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field
              label="Password"
              type="password"
              style="padding-left: 25%; padding-right: 25%"
              variant="underlined"
              :rules="[rules.required, rules.minLen]"
              v-model="validation.password"
              @keyup.enter="login"
            ></v-text-field>
          </v-col>
          <v-col>
            <router-link style="font-size: 15px" :to="{ path: '/Register' }"
              >Not Registered? Register here</router-link
            >
          </v-col>
          <v-divider length="800px"></v-divider>
          <v-col>
            <div class="text-right">
              <v-card-actions>
                <v-btn rounded @click="login">Login</v-btn>

                <v-btn color="red" rounded @click="reset">reset</v-btn>
              </v-card-actions>
            </div>
          </v-col>
        </v-row>
      </v-form>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, reactive, ref } from "vue";

export default defineComponent({
  name: "LoginDefault",

  data() {
    const rules = {
      required: (value: string) => !!value || "Required",
      minLen: (v: string) => v.length >= 8 || "Min. 8 Characters",
    };

    const validation = reactive({
      username: "",
      password: "",
    });

    const login = async () => {
      let registerRequest = ref(false);

      if (validation.username !== "" && validation.password.length >= 8) {
        try {
          registerRequest.value = await loginService.login(
            validation.username,
            validation.password
          );
          if (registerRequest.value)
            this.$router.push({ name: "home", query: { redirect: "/" } });
          else alert("something went wrong. Check your password and username.");
        } catch (e) {
          alert("something went wrong. Check your password and username.");
        }
        console.log(registerRequest.value);
      } else {
        alert("Something went wrong. Wrong username or password");
      }
    };

    const reset = () => {
      const obj = this.$refs.valid as any;
      if (obj !== undefined) obj.reset();
    };
    return { rules, validation, login, reset };
  },
});
</script>
