package main

// ExerciseDetail represents a single exercise with description and steps.
type ExerciseDetail struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Steps       []string `json:"steps"`
}

// ExerciseMap stores all muscles and their exercises.
var ExerciseMap = map[string][]ExerciseDetail{
	// --- Chest ---
	"chest": {
		{
			Name:        "Push-up",
			Description: "A bodyweight movement that targets the chest, shoulders, and triceps while engaging the core.",
			Steps: []string{
				"Place your hands firmly on the ground, directly under your shoulders.",
				"Flatten your back so your entire body forms a straight line from head to heels.",
				"Draw your shoulder blades back and down, keeping elbows close to your body as you lower yourself.",
				"Exhale as you push back up to the starting position.",
			},
		},
		{
			Name:        "Bench Press (Barbell)",
			Description: "A compound lift that targets the pectorals, triceps, and shoulders.",
			Steps: []string{
				"Lie on a flat bench with your feet flat on the floor.",
				"Grip the bar slightly wider than shoulder width.",
				"Lower the bar slowly to your mid-chest.",
				"Press the bar back up until your arms are fully extended.",
			},
		},
		{
			Name:        "Incline Dumbbell Press",
			Description: "A chest variation emphasizing the upper pectorals using dumbbells.",
			Steps: []string{
				"Set an incline bench to about 30–45 degrees.",
				"Hold a dumbbell in each hand at chest level.",
				"Press the weights upward until arms are fully extended.",
				"Lower them slowly back down and repeat.",
			},
		},
		{
			Name:        "Chest Press (Machine)",
			Description: "Machine-based exercise focusing on controlled chest activation.",
			Steps: []string{
				"Sit on the machine with the handles at chest level.",
				"Push the handles forward until arms are extended.",
				"Return slowly to the starting position.",
			},
		},
		{
			Name:        "Cable Crossover",
			Description: "Uses cable pulleys to engage inner and outer chest fibers.",
			Steps: []string{
				"Set both pulleys above shoulder height.",
				"Grab handles and step forward slightly.",
				"Bring your hands together in front of you in a hugging motion.",
				"Return slowly with control.",
			},
		},
	},

	// --- Back ---
	"back": {
		{
			Name:        "Pull-up",
			Description: "A bodyweight exercise that strengthens the back and biceps.",
			Steps: []string{
				"Grab a pull-up bar with palms facing away, hands shoulder-width apart.",
				"Engage your back and pull your chest toward the bar.",
				"Lower yourself in a controlled motion until arms are straight, then repeat.",
			},
		},
		{
			Name:        "Lat Pulldown (Machine)",
			Description: "Cable machine version of pull-up for lat isolation.",
			Steps: []string{
				"Sit with thighs under pads and grip bar wide.",
				"Pull the bar down to your upper chest.",
				"Control the bar back to starting position.",
			},
		},
		{
			Name:        "Seated Cable Row",
			Description: "Engages the middle back and rhomboids.",
			Steps: []string{
				"Sit on the machine and grab the handle.",
				"Pull toward your abdomen while keeping your chest upright.",
				"Slowly extend your arms back to the start.",
			},
		},
		{
			Name:        "Deadlift",
			Description: "Full-body compound lift that primarily targets the back and legs.",
			Steps: []string{
				"Stand with feet shoulder width and bar over midfoot.",
				"Grip the bar just outside your knees.",
				"Lift by driving through your heels, keeping your back flat.",
				"Lower with control.",
			},
		},
		{
			Name:        "T-Bar Row",
			Description: "A compound back exercise that targets the middle and lower lats.",
			Steps: []string{
				"Stand over the T-bar and grip the handles.",
				"Pull the weight toward your torso while keeping your back straight.",
				"Lower the weight slowly back to the starting position.",
			},
		},
	},

	// --- Legs ---
	"legs": {
		{
			Name:        "Squat (Barbell)",
			Description: "Core compound movement that targets quads, glutes, and hamstrings.",
			Steps: []string{
				"Stand with barbell across upper back.",
				"Bend knees and hips to lower down.",
				"Push through heels to return to standing.",
			},
		},
		{
			Name:        "Leg Press (Machine)",
			Description: "Machine-based compound leg exercise.",
			Steps: []string{
				"Sit on machine with feet shoulder-width on the platform.",
				"Lower the platform slowly until knees are 90 degrees.",
				"Push back to the starting position.",
			},
		},
		{
			Name:        "Hamstring Curl (Machine)",
			Description: "Isolates the hamstrings for strengthening and shaping.",
			Steps: []string{
				"Sit or lie on the machine and hook heels under the pad.",
				"Curl your legs back as far as possible.",
				"Return slowly to the starting position.",
			},
		},
		{
			Name:        "Leg Extension (Machine)",
			Description: "Targets quadriceps using resistance on the lower legs.",
			Steps: []string{
				"Sit on the machine and place shins behind the pad.",
				"Extend your legs until straight.",
				"Lower slowly back down.",
			},
		},
		{
			Name:        "Walking Lunge",
			Description: "A dynamic version of the lunge that improves balance and coordination.",
			Steps: []string{
				"Step forward with one leg, lowering both knees to 90 degrees.",
				"Push through the front heel to bring your back leg forward into the next step.",
				"Continue alternating legs in a walking motion.",
			},
		},
	},

	// --- Shoulders ---
	"shoulders": {
		{
			Name:        "Overhead Press (Dumbbell)",
			Description: "Develops overall shoulder strength and stability.",
			Steps: []string{
				"Hold dumbbells at shoulder level, palms forward.",
				"Press overhead until arms are extended.",
				"Lower slowly to the starting position.",
			},
		},
		{
			Name:        "Lateral Raise",
			Description: "Isolates the lateral deltoid for width.",
			Steps: []string{
				"Hold dumbbells at your sides.",
				"Raise arms out to the sides until parallel to the floor.",
				"Lower back down slowly.",
			},
		},
		{
			Name:        "Rear Delt Fly (Machine)",
			Description: "Targets the rear deltoids and upper back.",
			Steps: []string{
				"Sit facing the machine and grab handles.",
				"Pull handles outward until arms are straight back.",
				"Return with control.",
			},
		},
		{
			Name:        "Arnold Press",
			Description: "A shoulder press variation that emphasizes all three deltoid heads.",
			Steps: []string{
				"Hold dumbbells in front of your shoulders, palms facing you.",
				"Rotate palms outward while pressing overhead.",
				"Reverse the motion to return to the start.",
			},
		},
	},

	// --- Arms ---
	"arms": {
		{
			Name:        "Bicep Curl (Dumbbell)",
			Description: "Classic movement for bicep isolation.",
			Steps: []string{
				"Hold dumbbells by your sides with palms forward.",
				"Curl both arms up while keeping elbows still.",
				"Lower slowly to starting position.",
			},
		},
		{
			Name:        "Cable Curl (Machine)",
			Description: "Provides constant tension throughout the curl.",
			Steps: []string{
				"Stand facing the cable machine.",
				"Grip the bar attachment with palms up.",
				"Curl up and lower under control.",
			},
		},
		{
			Name:        "Tricep Pushdown (Cable Machine)",
			Description: "Targets the triceps using a cable attachment.",
			Steps: []string{
				"Grip the bar or rope at chest level.",
				"Push down until arms are extended.",
				"Return slowly to start.",
			},
		},
		{
			Name:        "Dips",
			Description: "Bodyweight exercise for triceps and chest.",
			Steps: []string{
				"Grip parallel bars and lower your body by bending elbows.",
				"Push back up to the starting position.",
			},
		},
	},

	// --- Abs ---
	"abs": {
		{
			Name:        "Cable Crunch (Machine)",
			Description: "Weighted ab exercise focusing on upper abs.",
			Steps: []string{
				"Attach a rope to the high pulley.",
				"Kneel down and hold rope by your head.",
				"Crunch down, contracting your abs.",
			},
		},
		{
			Name:        "Hanging Leg Raise",
			Description: "Targets lower abs and hip flexors.",
			Steps: []string{
				"Hang from a pull-up bar.",
				"Raise your legs until parallel to the floor.",
				"Lower slowly back down.",
			},
		},
		{
			Name:        "Ab Wheel Rollout",
			Description: "A dynamic core exercise that works abs, shoulders, and lower back.",
			Steps: []string{
				"Kneel on the floor with an ab wheel in front of you.",
				"Roll forward slowly until your body is nearly straight.",
				"Use your abs to pull back to the starting position.",
			},
		},
	},

	// --- Glutes ---
	"glutes": {
		{
			Name:        "Hip Thrust (Barbell)",
			Description: "A glute isolation exercise performed with a barbell.",
			Steps: []string{
				"Sit with your upper back against a bench, knees bent, feet flat.",
				"Place a barbell over your hips.",
				"Push through your heels to lift hips until your body forms a straight line.",
				"Squeeze glutes at the top, then lower hips back down.",
			},
		},
		{
			Name:        "Glute Kickback (Cable)",
			Description: "Targets the glutes using a cable attachment.",
			Steps: []string{
				"Attach an ankle strap to a low pulley cable.",
				"Extend your leg back while keeping a slight bend in the knee.",
				"Return to the start with control and repeat.",
			},
		},
	},

	// --- Calves ---
	"calves": {
		{
			Name:        "Standing Calf Raise",
			Description: "A simple exercise that strengthens the gastrocnemius and soleus.",
			Steps: []string{
				"Stand with your feet shoulder-width apart.",
				"Raise your heels as high as possible, then lower slowly.",
			},
		},
		{
			Name:        "Seated Calf Raise (Machine)",
			Description: "A seated variation focusing on the soleus muscle.",
			Steps: []string{
				"Sit with the balls of your feet on the platform and weight over your thighs.",
				"Raise your heels as high as possible, then lower under control.",
			},
		},
	},
	// --- Triceps ---
	"triceps": {
		{
			Name:        "Tricep Dips",
			Description: "Bodyweight exercise that targets triceps and shoulders.",
			Steps: []string{
				"Grip parallel bars with arms straight.",
				"Lower your body by bending elbows to about 90 degrees.",
				"Push back up until arms are extended.",
			},
		},
		{
			Name:        "Overhead Tricep Extension (Dumbbell)",
			Description: "Isolates the long head of the triceps.",
			Steps: []string{
				"Hold one dumbbell overhead with both hands.",
				"Lower the weight behind your head by bending elbows.",
				"Extend arms back to the top.",
			},
		},
	},

	// --- Biceps ---
	"biceps": {
		{
			Name:        "Barbell Curl",
			Description: "Classic bicep exercise building strength and size.",
			Steps: []string{
				"Hold a barbell with palms facing up, arms shoulder-width apart.",
				"Curl bar up while keeping elbows close to sides.",
				"Lower slowly to starting position.",
			},
		},
		{
			Name:        "Hammer Curl (Dumbbell)",
			Description: "Targets biceps and forearms with a neutral grip.",
			Steps: []string{
				"Hold dumbbells with palms facing inward.",
				"Curl weights up together.",
				"Lower under control.",
			},
		},
	},

	// --- Forearms ---
	"forearms": {
		{
			Name:        "Wrist Curl (Barbell)",
			Description: "Strengthens the forearm flexors.",
			Steps: []string{
				"Sit and rest forearms on your thighs holding a barbell.",
				"Let wrists extend downward, then curl bar up.",
				"Lower slowly and repeat.",
			},
		},
		{
			Name:        "Reverse Curl",
			Description: "Targets the brachioradialis and forearms.",
			Steps: []string{
				"Hold barbell with palms down (overhand grip).",
				"Curl bar toward shoulders, keeping elbows in place.",
				"Lower with control.",
			},
		},
	},

	// --- Traps ---
	"traps": {
		{
			Name:        "Barbell Shrug",
			Description: "Isolates upper traps for strength and posture.",
			Steps: []string{
				"Hold a barbell with arms fully extended, palms facing body.",
				"Lift shoulders as high as possible, hold briefly.",
				"Lower slowly to the start.",
			},
		},
		{
			Name:        "Dumbbell Upright Row",
			Description: "Builds traps and shoulders.",
			Steps: []string{
				"Hold dumbbells in front of thighs, palms facing body.",
				"Pull weights straight up toward chin, elbows out.",
				"Lower slowly to starting position.",
			},
		},
	},

	// --- Neck ---
	"neck": {
		{
			Name:        "Neck Flexion",
			Description: "Strengthens neck flexor muscles for posture.",
			Steps: []string{
				"Lie on your back with head hanging off bench.",
				"Lift your head toward chest, hold briefly.",
				"Lower slowly back down.",
			},
		},
		{
			Name:        "Neck Extension",
			Description: "Targets back of the neck for support and stability.",
			Steps: []string{
				"Lie face down on a bench with head hanging off.",
				"Lift head up toward ceiling, hold briefly.",
				"Lower slowly back down.",
			},
		},
	},

	// --- Obliques ---
	"obliques": {
		{
			Name:        "Russian Twist",
			Description: "Works obliques through rotational movement.",
			Steps: []string{
				"Sit on the ground with knees bent and lean back slightly.",
				"Hold a weight or ball and twist torso side to side.",
			},
		},
		{
			Name:        "Side Plank",
			Description: "Isometric exercise targeting obliques and core.",
			Steps: []string{
				"Lie on your side and support body on one elbow and feet.",
				"Raise hips to form a straight line from shoulders to feet.",
				"Hold position for time, then switch sides.",
			},
		},
	},

	// --- Quads ---
	"quads": {
		{
			Name:        "Front Squat (Barbell)",
			Description: "Emphasizes the quadriceps more than traditional squats.",
			Steps: []string{
				"Hold barbell across front shoulders with elbows forward.",
				"Lower into squat keeping chest upright.",
				"Push through heels to return up.",
			},
		},
		{
			Name:        "Leg Extension (Machine)",
			Description: "Isolates the quadriceps muscles.",
			Steps: []string{
				"Sit on the machine and place shins behind the pad.",
				"Extend your legs until straight.",
				"Lower slowly back down.",
			},
		},
	},

	// --- Hamstrings ---
	"hamstrings": {
		{
			Name:        "Romanian Deadlift",
			Description: "Strengthens hamstrings and glutes while improving flexibility.",
			Steps: []string{
				"Hold barbell at hip level, feet shoulder-width apart.",
				"Hinge at hips and lower bar while keeping back flat.",
				"Drive hips forward to return to starting position.",
			},
		},
		{
			Name:        "Lying Leg Curl (Machine)",
			Description: "Targets hamstrings with controlled isolation movement.",
			Steps: []string{
				"Lie face down on the machine and hook heels under the pad.",
				"Curl your legs toward your glutes.",
				"Lower slowly back to the start.",
			},
		},
	},
	// --- Abductors ---
	"abductors": {
		{
			Name:        "Cable Hip Abduction",
			Description: "Targets outer thighs and glute medius using cables.",
			Steps: []string{
				"Attach an ankle strap to a low pulley cable machine.",
				"Stand sideways to the machine and lift your outer leg away from your body.",
				"Slowly return to the start and repeat.",
			},
		},
		{
			Name:        "Seated Abduction (Machine)",
			Description: "Isolates hip abductor muscles for strength and tone.",
			Steps: []string{
				"Sit on the abduction machine with pads against your outer thighs.",
				"Push your legs apart as far as comfortable.",
				"Return slowly to the starting position.",
			},
		},
	},

	// --- Adductors ---
	"adductors": {
		{
			Name:        "Seated Adduction (Machine)",
			Description: "Targets the inner thighs for stability and leg control.",
			Steps: []string{
				"Sit on the adduction machine with pads on inner thighs.",
				"Bring legs together against resistance.",
				"Return slowly to starting position.",
			},
		},
		{
			Name:        "Side-Lying Leg Adduction",
			Description: "Bodyweight exercise for strengthening the inner thighs.",
			Steps: []string{
				"Lie on your side with lower leg straight and upper leg bent over it.",
				"Lift the lower leg upward, pause briefly, then lower back down.",
			},
		},
	},

	// --- Lats ---
	"lats": {
		{
			Name:        "Pull-up (Wide Grip)",
			Description: "Emphasizes lat development for back width.",
			Steps: []string{
				"Grab the bar wider than shoulder-width with palms facing away.",
				"Pull your chin up above the bar by engaging your lats.",
				"Lower slowly until arms are fully extended.",
			},
		},
		{
			Name:        "Straight-Arm Pulldown (Cable)",
			Description: "Isolates the lats through a full range of motion.",
			Steps: []string{
				"Stand facing a cable machine with straight arms gripping a bar.",
				"Pull bar down to your thighs while keeping arms straight.",
				"Return slowly to starting position.",
			},
		},
	},

	// --- Delts ---
	"delts": {
		{
			Name:        "Front Raise (Dumbbell)",
			Description: "Targets the anterior deltoids for shoulder definition.",
			Steps: []string{
				"Hold dumbbells in front of your thighs with palms facing down.",
				"Raise one or both arms to shoulder height.",
				"Lower under control and repeat.",
			},
		},
		{
			Name:        "Reverse Pec Deck",
			Description: "Strengthens rear deltoids for balanced shoulder development.",
			Steps: []string{
				"Sit facing the pec deck machine and grasp the handles.",
				"Pull arms back in a wide arc until in line with shoulders.",
				"Return slowly to the starting position.",
			},
		},
	},

	// --- Core ---
	"core": {
		{
			Name:        "Plank",
			Description: "Fundamental core stability exercise that builds endurance.",
			Steps: []string{
				"Place forearms on ground with elbows under shoulders.",
				"Keep body straight from head to heels and hold position.",
			},
		},
		{
			Name:        "Mountain Climbers",
			Description: "Dynamic full-body core exercise that improves cardio and coordination.",
			Steps: []string{
				"Start in a high plank position.",
				"Alternate driving knees toward your chest in a running motion.",
			},
		},
	},

	// --- Lower Back ---
	"lower_back": {
		{
			Name:        "Back Extension (Roman Chair)",
			Description: "Strengthens the lower back and spinal erectors.",
			Steps: []string{
				"Position yourself on the Roman chair with hips just above the pad.",
				"Lower your torso toward the floor, then raise until in line with legs.",
				"Keep movements slow and controlled.",
			},
		},
		{
			Name:        "Superman",
			Description: "Bodyweight exercise targeting lower back and glutes.",
			Steps: []string{
				"Lie face down on the floor with arms extended forward.",
				"Lift arms, chest, and legs simultaneously, pause briefly.",
				"Lower slowly to the floor.",
			},
		},
	},
}
