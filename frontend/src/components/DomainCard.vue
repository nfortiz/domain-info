<template>
    <b-card 
        :title="item.title || item.domain"
        :sub-title="item.domain"
        class="mt-5"
        >        
        <b-card-text>
           <img v-if="item.logo" :src="item.logo" style="height: "/> {{ item.title }}  
        </b-card-text>
        <b-card-footer v-if="disabled">
            Calificación de SSLLabs: <b-badge :variant="getVariant()"> {{ item.ssl_grade }} </b-badge>
        </b-card-footer>
        <b-button 
            v-if="!disabled" 
            @click="show(item.domain)" 
            variant="secondary"            
        >Detalles</b-button>
    </b-card>
</template>
<script>
import getGrade from '../utils/getGrade.js'
export default {
    name: 'DomainCard',
    props: {
        item: Object,
        disabled: Boolean
    },
    methods: {
        show: function(domain) {
            this.$router.push({ name: 'details', params: { domainName: domain }})
        },
        getVariant: function() {
            return getGrade(this.item.ssl_grade)
        }
    }
}
</script>
<style>
    img {
        height: 16px;
        width: auto;
    }
</style>