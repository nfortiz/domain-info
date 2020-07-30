<template>
    <b-card 
        :title="item.address"
        :sub-title="item.owner"
        class="mt-2"
        bg-variant="dark"
        text-variant="white"
    >        
        <b-card-text>
         Calificación del Servidor: <b-badge :variant="getVariant()"> {{ item.ssl_grade }} </b-badge>  
        </b-card-text>
    </b-card>
</template>
<script>
import getGrade from '../utils/getGrade.js'

export default {
    name: 'ServerCard',
    props: {
        item: Object
    },
    methods: {
        getVariant: function() {
            return getGrade(this.item.ssl_grade)
        }
    }
}
</script>