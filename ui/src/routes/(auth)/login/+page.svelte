<script lang="ts">
	import { goto } from '$app/navigation';

	let loading = false;

	function handleSubmit() {
		loading = true;

		setTimeout(() => {
			goto('/');
		}, 500);
	}
</script>

<div class="grid place-content-center h-3/4">
	<h1 class="text-5xl font-bold mb-12 text-center text-white">ChargePI</h1>

	<form
		class="card w-96 bg-base-200 shadow-xl flex flex-col space-y-4 items-center p-10"
		on:submit|preventDefault={handleSubmit}
	>
		<input
			name="username"
			type="text"
			placeholder="Username"
			class="input input-bordered w-full max-w-xs"
			required
		/>

		<input
			name="password"
			type="password"
			placeholder="Password"
			class="input input-bordered w-full max-w-xs"
			required
		/>

		<button type="submit" class="btn btn-primary w-2/3" class:loading>Login</button>
	</form>
</div>
