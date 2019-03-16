<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      fixed
      app
    >
      <v-list dense>
        <router-link to="home">
          <v-list-tile>
              <v-list-tile-action>
                <v-icon>home</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>home</v-list-tile-title>
              </v-list-tile-content>
          </v-list-tile>
        </router-link>
        <router-link to="settings">
          <v-list-tile>
              <v-list-tile-action>
                <v-icon>settings</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>settings</v-list-tile-title>
              </v-list-tile-content>
          </v-list-tile>
        </router-link>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar color="#69F0AE" dark fixed app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>go-vue-typescript-electron</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <router-view name="main"/>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="#69F0AE" app>
      <span class="white--text"></span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';
  @Component({
    props: {
      source: String,
    },
  })
  export default class Main extends Vue {
    /** data */
    private drawer: string = '';
  }
</script>