<script lang="ts" context="module">
	import Btn from '$lib/Button.svelte';
	const backgrounds = [
		'default',
		'transparent',
		'accent',
		'success',
		'warning',
		'danger',
		'info',
		'white'
	];
	const colors = ['default', 'accent', 'success', 'warning', 'danger', 'info'];
</script>

<script lang="ts">
	/*
bg-accent-500 dark:bg-accent-600
bg-success-500 dark:bg-success-600
bg-warning-500 dark:bg-warning-600
bg-danger-500 dark:bg-danger-600
bg-info-500 dark:bg-info-600
*/
</script>

<template>
	<table class="mt-24">
		<tr>
			<td>color\bg</td>
			{#each backgrounds as bg}
				<td>{bg}</td>
			{/each}
		</tr>
		{#each colors as color}
			<tr>
				<td style="padding: 1rem;">{color}</td>
				{#each backgrounds as bg}
					<td
						style="padding: 1rem;"
						class={bg === 'white' ? `bg-${color}-500 dark:bg-${color}-600` : ''}
					>
						<Btn icon="cross" text="Button" {color} {bg} />
					</td>
				{/each}
			</tr>
		{/each}
	</table>
</template>
