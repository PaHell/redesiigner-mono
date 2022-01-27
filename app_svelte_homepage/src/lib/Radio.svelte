<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
</script>
<script lang="ts">
	export let options = [
		{}
	];
	export let selected;
	export let opened = false;
</script>

<template>
	<div class="app-select">
		<button>
			<Icon name="home"/>
			<p>{ selected }</p>
		</button>
		<menu>
			{#each options as option}
				<button>
					<p>{ option }</p>
				</button>
			{/each}
		</menu>
	</div>
</template>

<style lang="postcss">
	.app-select {
		@apply w-64 h-28 mt-64;
	}
</style>
