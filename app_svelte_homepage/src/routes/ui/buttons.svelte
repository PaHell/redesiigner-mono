<script lang="ts" context="module">
	import Btn from '$lib/Button.svelte';
	import { theme } from '$lib/../store';
	const backgrounds = [
		'default',
		'accent',
		'success',
		'warning',
		'danger',
		'info'
	];
	const colors = ['default', 'accent', 'success', 'warning', 'danger', 'info'];
</script>

<script lang="ts">
/*
bg-accent-500  dark:bg-accent-600  text-accent-500  dark:text-accent-400 
bg-success-500 dark:bg-success-600 text-success-500 dark:text-success-400
bg-warning-500 dark:bg-warning-600 text-warning-500 dark:text-warning-400
bg-danger-500  dark:bg-danger-600  text-danger-500  dark:text-danger-400 
bg-info-500    dark:bg-info-600    text-info-500    dark:text-info-400   
*/
	let allLoading = true;
	import Select from "$lib/Select.svelte";

</script>

<template>
	<div class="w-full p-4 mt-8 rounded bg-gray-50 dark:bg-gray-800 border-gray-300 dark:border-gray-700">
		<div class="app-bar-button justify-center">
			<p>Theme</p>
			<Btn icon="device" text="OS Pref"
				active={$theme === 'os'}
				on:click={() => theme.set('os')} />
			<Btn icon="device" text="Light"
				active={$theme === 'light'}
				on:click={() => theme.set('light')} />
			<Btn icon="device" text="Dark"
				active={$theme === 'dark'}
				on:click={() => theme.set('dark')} />
		</div>
		<div class="flex flex-wrap justify-center space-2">
			{#each colors as bg}
			<div class="p-2 border-b dark:border-gray-700">
				<p class={`px-4 text-center text-lg text-${bg}-500 dark:text-${bg}-400`}>{bg}</p>
				<div class="app-bar-button">
					<Btn icon="cross" text="Transparent" color={bg} bg="transparent" />
					<Btn icon="cross" text="Loading" loading={allLoading} color={bg} bg="transparent" />
					<Btn icon="cross" text="Active" active color={bg} bg="transparent" />
					<Btn icon="cross" text="Disabled" disabled color={bg} bg="transparent" />
				</div>
				<div class="app-bar-button">
					<Btn icon="cross" text="Background" {bg} />
					<Btn icon="cross" text="Loading" loading={allLoading} {bg} />
					<Btn icon="cross" text="Active" active {bg} />
					<Btn icon="cross" text="Disabled" disabled {bg} />
				</div>
				<div class="app-bar-button">
					<Btn icon="cross" text="Color" color={bg} />
					<Btn icon="cross" text="Loading" loading={allLoading} color={bg} />
					<Btn icon="cross" text="Active" active color={bg} />
					<Btn icon="cross" text="Disabled" disabled color={bg} />
				</div>
				{#if bg != 'default'}
				<div class="{ `app-bar-button rounded-lg bg-${bg}-500` }">
					<Btn icon="cross" text="White" color={bg} bg="white" />
					<Btn icon="cross" text="Loading" loading={allLoading} color={bg} bg="white" />
					<Btn icon="cross" text="Active" active color={bg} bg="white" />
					<Btn icon="cross" text="Disabled" disabled color={bg} bg="white" />
				</div>
				{/if}
			</div>
			{/each}
		</div>
			<!--
				<Select/>
			-->
	</div>
</template>
