<script context="module" lang="ts">
	import Btn from '$lib/Button.svelte';
</script>

<script lang="ts">
	export let items : string[] = [];
	export let active = 0;
</script>

<template>
	<div>
		<div class="dark flex space-x-1">
			{#each items as item, index}
			<Btn
				text={item}
				active={active === index}
				bg={active === index ? 'default' : 'transparent'}
				disableLoading
				on:click={() => active = index}
				/>
			{/each}
		</div>
	</div>
	<!-- Mobile menu, show/hide based on menu state. -->
	<div class="hidden" id="mobile-menu">
		<div class="px-2 pt-2 pb-3 space-y-1">
			<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
			<a
				href="#"
				class="bg-gray-900 text-white block px-3 py-2 rounded-md text-base font-medium"
				aria-current="page">Dashboard</a
			>

			<a
				href="#"
				class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
				>Team</a
			>

			<a
				href="#"
				class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
				>Projects</a
			>

			<a
				href="#"
				class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
				>Calendar</a
			>
		</div>
	</div>
</template>

<style global lang="postcss">
</style>
