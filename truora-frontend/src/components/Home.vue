<template>
  <div class="hello">
    <b-container>
      <h2>Inserte el dominio</h2>
      <b-row align-h="center">
        <b-col cols="5">
          <b-input
            id="domain-input"
            v-model="form.domain"
            type="text"
            require
            placeholder="Enter domain. www.example.com"
            size="lg"
          ></b-input>
        </b-col>
        <b-col cols="1">
          <b-button 
            type="button" 
            variant="primary" 
            @click="onSubmit"
            size="lg"
          >Buscar</b-button>
        </b-col>
      </b-row>
      <b-row>
        <b-col cols="6">
          <Domain 
            v-if="domain"
            :item="domain"
            class="mt-5"
          >        
          </Domain>
        </b-col>
        
        <b-col cols="4">
          <Domain
            v-for="(item, index) in history"
            :item="item"
            :key="index"
          > 
        </Domain>
        </b-col>
      </b-row>

    </b-container>
  </div>
</template>

<script>
import Domain from './DomainCard'
import { getDomain, getHistory } from '../utils/getData'

export default {
  name: 'Home',
  components: {
    Domain
  },
  props: {
    msg: String
  },
  data: function() {
    return {
      history: [],
      error: false,
      domain: null,
      form: {
        domain: '',
      }
    }
  },
  methods: {
    onSubmit: function() {
      getDomain(this.$data.form.domain)
        .then(data => { this.$data.domain = data })
    }
  },
  created: function() {
    getHistory()
      .then(data => { this.$data.history = data.items })
  },
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
</style>
