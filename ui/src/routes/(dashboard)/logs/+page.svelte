<script lang="ts">
	const logs = [
		{
			timestamp: new Date(),
			message: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit.',
		},
		{
			timestamp: new Date(),
			message: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit.',
		},
		{
			timestamp: new Date(),
			message: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit.',
		},
	];
</script>

<div class="overflow-x-auto">
	<table class="table">
		<thead>
			<tr>
				<th>Timestamp</th>
				<th>Message</th>
			</tr>
		</thead>
		<tbody>
			{#each logs as log}
				<tr>
					<td>{log.timestamp.toISOString()}</td>
					<td>{log.message}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
