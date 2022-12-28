<script lang="ts">
	import { faMoon, faSun } from '@fortawesome/free-solid-svg-icons';
	import { onMount } from 'svelte';
	import Fa from 'svelte-fa/src/fa.svelte';
	import { themeChange } from 'theme-change';

	let clazz: string = '';
	export { clazz as class };

	let theme: string | null = null;

	onMount(() => {
		themeChange(false);
		theme = document.documentElement.getAttribute('data-theme');
	});
</script>

<label class={`swap swap-rotate ${clazz}`}>
	<input
		type="checkbox"
		data-toggle-theme="dark,light"
		data-act-class="data-theme"
		checked={theme === 'light'}
	/>
	<Fa class="swap-on fill-current" icon={faSun} />
	<Fa class="swap-off fill-current" icon={faMoon} />
</label>
