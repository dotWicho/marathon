package marathon

// Metrics Marathon server metrics results
type Metrics struct {
	Version  string `json:"version"`
	Counters struct {
		MarathonMesosOfferOperationsLaunchGroupCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offer-operations.launch-group.counter"`
		MarathonMesosOffersIncomingCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offers.incoming.counter"`
		MarathonDebugMesosOffersSavingTasksErrorsCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.mesos.offers.saving-tasks-errors.counter"`
		MarathonMesosCallsReviveCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.calls.revive.counter"`
		MarathonDebugMesosOffersUnprocessableCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.mesos.offers.unprocessable.counter"`
		MarathonMesosTaskUpdatesTaskStartingCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-starting.counter"`
		MarathonMesosOffersUsedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offers.used.counter"`
		MarathonMesosOffersDeclinedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offers.declined.counter"`
		MarathonMesosOfferOperationsReserveCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offer-operations.reserve.counter"`
		MarathonDebugPersistenceCacheGetDeploymentHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.get.deployment.hit.counter"`
		MarathonMesosTaskUpdatesTaskKillingCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-killing.counter"`
		MarathonMesosTaskUpdatesTaskKilledCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-killed.counter"`
		MarathonDeploymentsCounter struct {
			Count int `json:"count"`
		} `json:"marathon.deployments.counter"`
		MarathonDebugPersistenceCacheIdsPodsHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.ids.pods.hit.counter"`
		MarathonHTTPResponsesEventStreamSizeCounterBytes struct {
			Count int `json:"count"`
		} `json:"marathon.http.responses.event-stream.size.counter.bytes"`
		MarathonMesosCallsSuppressCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.calls.suppress.counter"`
		MarathonDebugPersistenceCacheGetTaskFailuresHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.get.taskFailures.hit.counter"`
		MarathonMesosTaskUpdatesTaskFailedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-failed.counter"`
		MarathonDebugPersistenceCacheGetFrameworkIDHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.get.framework-id.hit.counter"`
		MarathonMesosTaskUpdatesTaskRunningCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-running.counter"`
		MarathonDebugPersistenceCacheGetAppsHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.get.apps.hit.counter"`
		MarathonMesosTaskUpdatesTaskFinishedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-finished.counter"`
		MarathonMesosTaskUpdatesTaskStagingCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.task-updates.task-staging.counter"`
		MarathonHTTPResponsesSizeCounterBytes struct {
			Count int64 `json:"count"`
		} `json:"marathon.http.responses.size.counter.bytes"`
		MarathonDebugPersistenceCacheIdsAppsHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.ids.apps.hit.counter"`
		MarathonHTTPResponsesSizeGzippedCounterBytes struct {
			Count int64 `json:"count"`
		} `json:"marathon.http.responses.size.gzipped.counter.bytes"`
		MarathonDeploymentsDismissedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.deployments.dismissed.counter"`
		MarathonDebugPersistenceCacheIdsDeploymentHitCounter struct {
			Count int `json:"count"`
		} `json:"marathon.debug.persistence.cache.ids.deployment.hit.counter"`
		MarathonTasksLaunchedCounter struct {
			Count int `json:"count"`
		} `json:"marathon.tasks.launched.counter"`
		MarathonMesosOfferOperationsLaunchCounter struct {
			Count int `json:"count"`
		} `json:"marathon.mesos.offer-operations.launch.counter"`
		MarathonPersistenceGcRunsCounter struct {
			Count int `json:"count"`
		} `json:"marathon.persistence.gc.runs.counter"`
		MarathonHTTPRequestsSizeCounterBytes struct {
			Count int `json:"count"`
		} `json:"marathon.http.requests.size.counter.bytes"`
	} `json:"counters"`
	Gauges struct {
		MarathonJvmMemoryPoolsCodeCacheMaxGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.code-cache.max.gauge.bytes"`
		MarathonJvmMemoryPoolsMetaspaceCommittedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.metaspace.committed.gauge.bytes"`
		MarathonDebugHTTPRequests4XxTo15MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.4xx-to-15m-rate-ratio.gauge"`
		MarathonJvmMemoryPoolsCodeCacheUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.code-cache.usage.gauge"`
		MarathonJvmMemoryTotalMaxGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.total.max.gauge.bytes"`
		MarathonJvmBuffersMappedMemoryUsedGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.buffers.mapped.memory.used.gauge.bytes"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceCommittedGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.committed.gauge.bytes"`
		MarathonJvmGcPsScavengeCollectionsGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.gc.ps-scavenge.collections.gauge"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceMaxGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.max.gauge.bytes"`
		MarathonInstancesInflightKillsGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.instances.inflight-kills.gauge"`
		MarathonJvmThreadsRunnableGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.runnable.gauge"`
		MarathonInstancesInflightKillAttemptsGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.instances.inflight-kill-attempts.gauge"`
		MarathonJvmMemoryPoolsPsOldGenUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.used.gauge.bytes"`
		MarathonDebugHTTPDispatchesActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.dispatches.active.gauge"`
		MarathonJvmMemoryNonHeapInitGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.non-heap.init.gauge.bytes"`
		MarathonJvmMemoryTotalUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.total.used.gauge.bytes"`
		MarathonJvmMemoryNonHeapUsageGauge struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.non-heap.usage.gauge"`
		MarathonJvmBuffersMappedCapacityGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.buffers.mapped.capacity.gauge.bytes"`
		MarathonJvmThreadsDaemonGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.daemon.gauge"`
		MarathonDebugHTTPRequestsSuspendedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.suspended.gauge"`
		MarathonDeploymentsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.deployments.active.gauge"`
		MarathonLeadershipDurationGaugeSeconds struct {
			Value float64 `json:"value"`
		} `json:"marathon.leadership.duration.gauge.seconds"`
		MarathonDebugHTTPRequests4XxTo5MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.4xx-to-5m-rate-ratio.gauge"`
		MarathonJvmMemoryPoolsPsEdenSpaceCommittedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.committed.gauge.bytes"`
		MarathonDebugHTTPRequests5XxTo5MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.5xx-to-5m-rate-ratio.gauge"`
		MarathonJvmMemoryNonHeapCommittedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.non-heap.committed.gauge.bytes"`
		MarathonJvmMemoryTotalInitGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.total.init.gauge.bytes"`
		MarathonJvmMemoryPoolsPsEdenSpaceInitGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.init.gauge.bytes"`
		MarathonJvmMemoryPoolsPsEdenSpaceUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.usage.gauge"`
		MarathonJvmMemoryPoolsCodeCacheCommittedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.code-cache.committed.gauge.bytes"`
		MarathonJvmMemoryHeapCommittedGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.heap.committed.gauge.bytes"`
		MarathonJvmMemoryPoolsMetaspaceUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.metaspace.used.gauge.bytes"`
		MarathonJvmThreadsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.active.gauge"`
		MarathonPodsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.pods.active.gauge"`
		MarathonJvmThreadsBlockedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.blocked.gauge"`
		MarathonJvmMemoryPoolsPsOldGenUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.usage.gauge"`
		MarathonJvmBuffersDirectMemoryUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.buffers.direct.memory.used.gauge.bytes"`
		MarathonJvmMemoryPoolsPsOldGenCommittedGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.committed.gauge.bytes"`
		MarathonUptimeGaugeSeconds struct {
			Value float64 `json:"value"`
		} `json:"marathon.uptime.gauge.seconds"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceUsedAfterGcGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.used-after-gc.gauge.bytes"`
		MarathonHTTPRequestsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.http.requests.active.gauge"`
		MarathonJvmMemoryPoolsCompressedClassSpaceCommittedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.compressed-class-space.committed.gauge.bytes"`
		MarathonJvmMemoryPoolsPsEdenSpaceUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.used.gauge.bytes"`
		MarathonJvmMemoryPoolsMetaspaceInitGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.metaspace.init.gauge.bytes"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceUsedGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.used.gauge.bytes"`
		MarathonDebugOfferMatcherQueueSizeGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.offer-matcher.queue.size.gauge"`
		MarathonJvmMemoryPoolsPsEdenSpaceUsedAfterGcGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.used-after-gc.gauge.bytes"`
		MarathonJvmMemoryHeapUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.heap.usage.gauge"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceInitGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.init.gauge.bytes"`
		MarathonJvmMemoryHeapMaxGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.heap.max.gauge.bytes"`
		MarathonHTTPEventStreamsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.http.event-streams.active.gauge"`
		MarathonJvmMemoryPoolsCompressedClassSpaceMaxGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.compressed-class-space.max.gauge.bytes"`
		MarathonJvmMemoryHeapUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.heap.used.gauge.bytes"`
		MarathonJvmGcPsMarksweepCollectionsDurationGaugeSeconds struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.gc.ps-marksweep.collections.duration.gauge.seconds"`
		MarathonJvmGcPsScavengeCollectionsDurationGaugeSeconds struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.gc.ps-scavenge.collections.duration.gauge.seconds"`
		MarathonJvmMemoryPoolsPsOldGenMaxGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.max.gauge.bytes"`
		MarathonJvmThreadsTimedWaitingGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.timed-waiting.gauge"`
		MarathonJvmThreadsTerminatedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.terminated.gauge"`
		MarathonJvmMemoryTotalCommittedGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.total.committed.gauge.bytes"`
		MarathonJvmMemoryPoolsPsEdenSpaceMaxGaugeBytes struct {
			Value int64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-eden-space.max.gauge.bytes"`
		MarathonJvmThreadsNewGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.new.gauge"`
		MarathonJvmMemoryNonHeapMaxGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.non-heap.max.gauge.bytes"`
		MarathonJvmMemoryPoolsPsSurvivorSpaceUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-survivor-space.usage.gauge"`
		MarathonJvmMemoryPoolsMetaspaceUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.metaspace.usage.gauge"`
		MarathonGroupsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.groups.active.gauge"`
		MarathonJvmMemoryPoolsCodeCacheUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.code-cache.used.gauge.bytes"`
		MarathonJvmMemoryPoolsPsOldGenInitGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.init.gauge.bytes"`
		MarathonDebugOfferMatcherTokensGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.offer-matcher.tokens.gauge"`
		MarathonDebugHTTPRequests5XxTo1MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.5xx-to-1m-rate-ratio.gauge"`
		MarathonDebugHTTPRequests4XxTo1MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.4xx-to-1m-rate-ratio.gauge"`
		MarathonJvmGcPsMarksweepCollectionsGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.gc.ps-marksweep.collections.gauge"`
		MarathonInstancesStagedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.instances.staged.gauge"`
		MarathonJvmMemoryPoolsMetaspaceMaxGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.metaspace.max.gauge.bytes"`
		MarathonJvmThreadsWaitingGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.waiting.gauge"`
		MarathonJvmMemoryPoolsCompressedClassSpaceUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.compressed-class-space.used.gauge.bytes"`
		MarathonDebugRootGroupUpdatesActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.root-group.updates.active.gauge"`
		MarathonJvmBuffersDirectGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.buffers.direct.gauge"`
		MarathonJvmMemoryPoolsCompressedClassSpaceUsageGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.compressed-class-space.usage.gauge"`
		MarathonJvmMemoryNonHeapUsedGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.non-heap.used.gauge.bytes"`
		MarathonJvmMemoryHeapInitGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.heap.init.gauge.bytes"`
		MarathonAppsActiveGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.apps.active.gauge"`
		MarathonInstancesLaunchOverdueGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.instances.launch-overdue.gauge"`
		MarathonInstancesRunningGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.instances.running.gauge"`
		MarathonJvmMemoryPoolsCodeCacheInitGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.code-cache.init.gauge.bytes"`
		MarathonJvmBuffersDirectCapacityGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.buffers.direct.capacity.gauge.bytes"`
		MarathonDebugHTTPRequests5XxTo15MRateRatioGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.debug.http.requests.5xx-to-15m-rate-ratio.gauge"`
		MarathonJvmMemoryPoolsPsOldGenUsedAfterGcGaugeBytes struct {
			Value int `json:"value"`
		} `json:"marathon.jvm.memory.pools.ps-old-gen.used-after-gc.gauge.bytes"`
		MarathonJvmThreadsDeadlockedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.threads.deadlocked.gauge"`
		MarathonJvmMemoryPoolsCompressedClassSpaceInitGaugeBytes struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.memory.pools.compressed-class-space.init.gauge.bytes"`
		MarathonJvmBuffersMappedGauge struct {
			Value float64 `json:"value"`
		} `json:"marathon.jvm.buffers.mapped.gauge"`
	} `json:"gauges"`
	Histograms struct {
	} `json:"histograms"`
	Meters struct {
		MarathonHTTPResponses5XxRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.http.responses.5xx.rate.meter"`
		MarathonHTTPResponses4XxRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.http.responses.4xx.rate.meter"`
		MarathonHTTPResponses2XxRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.http.responses.2xx.rate.meter"`
		MarathonDebugHTTPDispatchesAsyncRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.debug.http.dispatches.async.rate.meter"`
		MarathonHTTPResponses1XxRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.http.responses.1xx.rate.meter"`
		MarathonDebugHTTPDispatchesAsyncTimeoutsRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.debug.http.dispatches.async.timeouts.rate.meter"`
		MarathonHTTPResponses3XxRateMeter struct {
			Count    int     `json:"count"`
			M1Rate   float64 `json:"m1_rate"`
			M5Rate   float64 `json:"m5_rate"`
			M15Rate  float64 `json:"m15_rate"`
			MeanRate float64 `json:"mean_rate"`
			Units    string  `json:"units"`
		} `json:"marathon.http.responses.3xx.rate.meter"`
	} `json:"meters"`
	Timers struct {
		MarathonDebugPersistenceOperationsStoreDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.persistence.operations.store.duration.timer.seconds"`
		MarathonDebugInstanceTrackerUpdateStepsPostTaskStatusEventDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.update-steps.post-task-status-event.duration.timer.seconds"`
		MarathonHTTPRequestsMoveDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.move.duration.timer.seconds"`
		MarathonDebugPersistenceOperationsVersionsDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.persistence.operations.versions.duration.timer.seconds"`
		MarathonDebugHTTPDispatchesDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.http.dispatches.duration.timer.seconds"`
		MarathonDebugInstanceTrackerResolveTasksByAppDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.resolve-tasks-by-app-duration.timer.seconds"`
		MarathonHTTPRequestsOtherDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.other.duration.timer.seconds"`
		MarathonDebugInstanceTrackerUpdateStepsNotifyHealthCheckManagerDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.update-steps.notify-health-check-manager.duration.timer.seconds"`
		MarathonPersistenceGcCompactionDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.persistence.gc.compaction.duration.timer.seconds"`
		MarathonDebugInstanceTrackerUpdateStepsScaleAppDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.update-steps.scale-app.duration.timer.seconds"`
		MarathonHTTPRequestsGetDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.get.duration.timer.seconds"`
		MarathonDebugCurrentLeaderRetrievalDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.current-leader.retrieval.duration.timer.seconds"`
		MarathonDebugInstanceTrackerUpdateStepsNotifyRateLimiterDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.update-steps.notify-rate-limiter.duration.timer.seconds"`
		MarathonDebugMesosOffersSavingTasksDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.mesos.offers.saving-tasks-duration.timer.seconds"`
		MarathonHTTPRequestsPutDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.put.duration.timer.seconds"`
		MarathonDebugKillingUnknownTaskDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.killing-unknown-task-duration.timer.seconds"`
		MarathonPersistenceGcScanDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.persistence.gc.scan.duration.timer.seconds"`
		MarathonHTTPRequestsOptionsDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.options.duration.timer.seconds"`
		MarathonHTTPRequestsDeleteDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.delete.duration.timer.seconds"`
		MarathonHTTPRequestsHeadDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.head.duration.timer.seconds"`
		MarathonDebugPersistenceOperationsGetDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.persistence.operations.get.duration.timer.seconds"`
		MarathonHTTPRequestsPostDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.post.duration.timer.seconds"`
		MarathonDebugInstanceTrackerUpdateStepsNotifyLaunchQueueDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.instance-tracker.update-steps.notify-launch-queue.duration.timer.seconds"`
		MarathonDebugPublishingTaskStatusUpdateDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.publishing-task-status-update-duration.timer.seconds"`
		MarathonDebugPersistenceOperationsDeleteDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.persistence.operations.delete.duration.timer.seconds"`
		MarathonDebugPersistenceOperationsIdsDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.persistence.operations.ids.duration.timer.seconds"`
		MarathonHTTPRequestsTraceDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.trace.duration.timer.seconds"`
		MarathonHTTPRequestsDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.duration.timer.seconds"`
		MarathonHTTPRequestsConnectDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.http.requests.connect.duration.timer.seconds"`
		MarathonDebugMesosOffersMatchingDurationTimerSeconds struct {
			Count         int     `json:"count"`
			Min           float64 `json:"min"`
			Mean          float64 `json:"mean"`
			Max           float64 `json:"max"`
			P50           float64 `json:"p50"`
			P75           float64 `json:"p75"`
			P95           float64 `json:"p95"`
			P98           float64 `json:"p98"`
			P99           float64 `json:"p99"`
			P999          float64 `json:"p999"`
			Stddev        float64 `json:"stddev"`
			M1Rate        float64 `json:"m1_rate"`
			M5Rate        float64 `json:"m5_rate"`
			M15Rate       float64 `json:"m15_rate"`
			MeanRate      float64 `json:"mean_rate"`
			DurationUnits string  `json:"duration_units"`
			RateUnits     string  `json:"rate_units"`
		} `json:"marathon.debug.mesos.offers.matching-duration.timer.seconds"`
	} `json:"timers"`
}
