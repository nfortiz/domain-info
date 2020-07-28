<template>
    <b-container>
        <b-row v-if="item" align-h="center">
            <b-col cols="6" v-if="item.servers_changed">
                <b-alert show variant="warning">Los servidores han cambiado, anterior calificación {{ item.previous_ssl_grade }} </b-alert>
            </b-col> 
        </b-row>
        <b-row v-if="item" align-h="center">
            <b-col cols="6">
                <DomainCard :item="item" disabled></DomainCard>
            </b-col>            
        </b-row>
        <b-row v-if="item" class="my-5">
            <b-col cols="12">
                <h4>Servers</h4>

            </b-col>
            <b-col cols="4" v-for="(server, idx) in item.servers" :key="idx">
                <ServerCard :item="server"></ServerCard>
            </b-col>
        </b-row>
    </b-container>
</template>
<script>
import DomainCard from './DomainCard'
import ServerCard from './ServerCard'
import { getDomain } from '../utils/getData'

export default {
    name: 'Details',
    components: {
        DomainCard,
        ServerCard
    },
    data: function () {
        return {
            item: null
        }
    },
    created: function() {
        getDomain(this.$route.params.domainName)
            .then(data => { this.item = data })
    },
}
</script>