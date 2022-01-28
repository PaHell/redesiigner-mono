<script lang="ts" context="module">
	import Btn from '$lib/Button.svelte';
	import Sidebar from '$lib/Sidebar.svelte';
	import ProfileListEntry from '$lib/ProfileListEntry.svelte';
	import { bannerStore, modalStore, notificationStore } from '../store';
	import Banner from '$lib/banner/Banner.svelte';

	const styles = ['default', 'accent', 'success', 'warning', 'danger', 'info'];
</script>

<script lang="ts">
	const backgrounds = ['default', 'transparent', 'accent', 'success', 'warning', 'danger', 'info'];
	const colors = ['default', 'accent', 'success', 'warning', 'danger', 'info'];
</script>

<template>
	<table>
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
            <td style="padding: 1rem;">
                <Btn icon="cross" text="Button" {color} {bg}/>
            </td>
            {/each}
        </tr>
        {/each}
    </table>
</template>
