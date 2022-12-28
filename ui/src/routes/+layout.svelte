<script lang="ts">
	import '../app.css';

	export const ssr = false;

	const ocppConnected = false;
	const version = '0.1';
</script>

<div class="flex flex-col h-screen">
	<div class="h-full">
		<slot />
	</div>

	<footer class="footer items-center p-4 bg-neutral text-neutral-content">
		<div class="items-center grid-flow-col">
			<p>Copyright ChargePI © 2022 - All right reserved</p>
		</div>
		<div class="grid-flow-col gap-2 md:place-self-center md:justify-self-end">
			<div
				class="badge badge-success"
				class:badge-success={ocppConnected}
				class:badge-error={!ocppConnected}
			>
				{ocppConnected ? 'Connected' : 'Not Connected'}
			</div>
			<div class="badge badge-info">v{version}</div>
		</div>
	</footer>
</div>
