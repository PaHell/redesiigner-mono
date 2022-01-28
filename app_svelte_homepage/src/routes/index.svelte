<script lang="ts" context="module">
	import Radio from '$lib/Radio.svelte';
	import Select from '$lib/Select.svelte';
	import List from '$lib/CardPagination.svelte';
	import InfoCard from '$lib/InfoCard.svelte';
	import Table from '$lib/Table.svelte';
	import CardHeader from '$lib/CardHeader.svelte';
	import Login from '$lib/forms/Login.svelte';
</script>

<script lang="ts">
</script>

<template>
	<div class="w-full h-full flex justify-center">
		<div class="w-1/2">
			<Radio />
			<Login />
			<List />
			<Table />
			<Select />
			<CardHeader />
			<InfoCard />
		</div>
	</div>
</template>
