# Context Based Identity-Management Method

**Abstract**
The idea proposes a new method for mobile and immobile asset protection and secure life cycle management using network assisted and network provided events.
The method to
1. Generate mobility and communication context events.
2. Store communication context event data.
3. Generate a Hash Tag based on stored event data.
4. Authorize group users to acquire hash tags for multiuser authentication.
5. Induced/Virtual roaming for relative events.
6. Transferring the Assets to the next trusted user using mobile network events 

Events can be location events or device state events or mobile context events specific to device or relative to device used to prove the authorized owner 
are specific to these domain 

1.	IoT: Smart City
2.	General Public: Protecting/Monitoring the assets 
3.	Government/Private Body: Protecting/Monitoring the assets
4.	Context Based Resource Authorization and Authentication
5.	Context Based Identity Management.
6.	Supply Chain Controls
7.	Electronic tokens for group asset managments monitoring and secure transfer

Some mobility Event Examples
1.	Location Events (Real, Induced,Relative) + UE driven 
2.	UE status Event (Real, Induced) + UE driven 
3.	Roaming Events (Real, Induced) (includes space roaming)
4.	Re-Attach Event (Real, Induced)
5.	Detach Event (Induced)
6.	Relative Attach Event (Real)
7.	OAM events due to UE service usage or any counters related to UE-Activity like BRM counters [ CPU,Mem,disk,transaction contexts,connections etc.,]
8.	Charging Events due to UE service usage
9.	Attach from type of network event
10.	Attach from specified Device event
11.	Receive SMS event(Real, Induced,Relative) 
12.	Send SMS event (Real, Induced,Relative) 
13.	Subscribe for SMS event.(Real, Induced,Relative) 
14.	Generate email event(Real, Induced,Relative) 
15.	Get Charging event.
16.	Handover from RAT (inter/intra) events (Real, Induced,Relative) 
17.	Handover from different access events(includes satellite access)(Real, Induced,Relative) 
18.	Using multiple SIMs from multi-vendor operator related events(Real, Induced,Relative) 
19.	Events due to OTT(Real, Induced,Relative) 
20.	Change Device ownership event(Real, Induced,Relative) 
21.	Nominate the leader event in case of multisim based on schedules(Real, Induced,Relative) 
22.	UE activity from specific pheriperal devices ( Sensors,monitors ,cameras)(Real, Induced,Relative) 
23.	UE activity from specific version of trusted applicaiton (Real, Induced,Relative) 
24.	UE activity based on device age,physical state(Real, Induced,Relative) 
25.	Running mobile app residing in specific geographic location in cloud (Real, Induced,Relative) 
26.	Device access events based on white/black list of biometric information A-Party,CUG
27.	Device access events based on white/black list of speech and guesture of A-Party,CUG
28.	Device access events based on relation of other measuring devices ( relative to temperature sensor, electricity consumption or user of CUG or any one in CUG,usage of    	 data on a day)
29.	XR ( AR/VR ) based events in control and data plane that can be (Real, Induced,Relative)  
30.	Attaching from specific network slice event in combination with above events
31.	Events due to self Activations of SIM and linking to primary/secondary devices including CUG
32.	Events induced due to alternative messaging mechanisms example from Digital twin UEs or DTIs or UEs.
33.	Any statistical based learning and predictions from OAM for detection and estimation of Mobility context events / OAM events using AI/ML techniques.
34.	NF / User Access from specific tenants / multi tenants in case of multi tenant capable CN Registers or multi tenant capable NFs and Service Registery and discovery
35.	User / Device proximity measurement events < especially in case of UE misuses >
36.	Proximity event information event and monitoring  via alternate devices , RATs, 3GPP and non 3GPP services.
	Example case may cover alos : when user is not in the contact of device or in proximity the device has AI enabled assistant application 
	who can trigger SMS to CUG or to send mail to mail box and such applciation can be authorized by network operator.
37.	Relative Events from secondary Device.
38.	Relative Challenge events from last "N" calls
39.	Relative Challenge events indicating personal tokens in messages and subscribing for the same UE originated control signalling
40.	All UE specific application context static subscription configuration in profile including agreed named IP and Ports or static APN/DNN
	used to activate knowledge base authorization challenges.
41.     UE's Digital Twin based authorization.
42. Core network operator managed whitelist and blacklist of DTIs or allowed DTIs driven from subscriber provisioning context including allowed IMSI, MSISDN,IMEI,ME-ID as DTI as pre-provisioning context.



**Realization**
 ![image](https://github.com/KiranCS-17/identity-managment/blob/main/figure-1.png)
 
  1. E1, E2: Location Events generated by Devices 
  2. S1, S2: Location Events generated by RANs w.r.t source.
  3. S3, S4: Location Event Generation control by cloud core
  4. S5: Event store with specific pattern.

  Vicinity-Vr is the case of roaming by the same device D1 of the CUG.
  
  ##  Sample Message Flow and descriptions
  ![image](https://github.com/KiranCS-17/identity-managment/blob/main/figure-2.png)
  
**Note:**
The method of generating the events, inducing it and maintaining the records of events for UE and group of UEs, is hosted inside an AF called Communication Context & Event  Management Function.( CCEF)
This AF is responsible for producing the Authorization events and creating challenge factors for fraud user using time stamp or any related events as a question.
Some of the concepts of induced virtual roaming with AR-VR facilitation to create mobile communication context events that are encoded in 5G Data stores like UDR/UDSF.
The new application function CCEM (communication context and event Management) is used to trigger mobility events related to asset protection or monitoring/communication patters detections and formulating Event data stored in granular time series DB in compact method with context data chains. 
The Fraud user will be challenged to prove the Mobility Context Events and its relative events, at specified granular time or chain of times that becomes NP hard problem to solve and with multiple hashes the trusted user can request for recursive challenges with the recursion depth configured (medium,high,or greedy) . 
Block chain method can be one implementation method to realize this concept.
Google APIs can be used as an example for Location based Authentication and Authorization apart from 3GPP defined procedures.
The CCEF can be co-located inside the Core Network elements like HSS/UDM/NWDAF/OAM or any 3GPP or non 3GPP NF/service providers  as a new function or  a standalone function

1.  UE/CUG performs authentication and attach subscribing for this new monitoring service.
2.  RAN forwards this request to Serving Nodes (MME, AMF).
3.  Serving node asks HSS+UDM to perform authentication authorization.
4.  HSS+UDM reads the subscriber profile in DB.
5.  HSS+UDM gets the subscriber profile response from DB.
    5.1,5.2,5.3 The UE/CUG is authenticated and attached and acknowledged by network.
6.  HSS/UDM trigger CCEF to perform the service execution for this UE/CUG.
7.  CCEF asks HSS to generate the location events, or any other events mentioned in(Type of Events).
8.  HSS/UDM queries the serving node to get the event data.
9.  HSS/UDM gets event data from serving nodes.
10. HSS/UDM forwards the event data to CCEF.
11. CCEF acknowledges the HSS with permutation data.
        At this step the CCEF implicitly subscribes for permutation data movement from UDR to UDSF based on time stamp.
12. HSS/UDM updates the permutation data into the subscriber profile/CUG.
13. The CCEF Executed the Algorithm A1 and stores the output into UDSF in time series DB.
14. The UDSF acknowledges the data.

Based on user query or periodicity the CCEF can produce the Geo Hash with Algorithm A2
and supply to user that can be used as challenge against fraud.

**Goal**: The goal of this project is to realize the CCEF Application Function as PoC in step 1.
          Further steps focuses on performance aspects concerning response times.

**Prerequisites** : The UE has trusted application downloaded by the security service provider, the application in UE also produces mobility events / contexts 
                the trused application in UE also receives messages from trusted network to trigger events to produce mobility events / contexts.

**Sample Real Usage Example**
The CCEF instructs the user to perform induced roaming using Virtual reality method.
The user uses Virtual reality headsets and viewers and tries to roam from source (either chosen by network or by user)
and moves to the destination (Geo Asset) using a route map.
This virtually creates a roaming scenario.
In this roaming scenario the specific events can be induced by operator between the resource and CUG and the owner  
by using CCEF to get the location information, user state, the user specific contexts.
The other relative events from control plane elements are recorded as event data stored in UDSF.
The event can be fully constructed based on the induced roaming. 

the CCEF can host the ML Training algorithm for prediction and pre/post detection of mobility context events.

A hash-code is generated for this with input as Event data and stored against the time stamp in UDSF.

**Algorithm A1:**
The compression uses encoding methods internally with below strategy.
1. Uses the profile reference from UDR / LDAP (basically UID of the user/CUG)
2. Maps the event and the correlation events with unique numbers
3. Uses the dictionary methods to compress the Event data received from serving nodes or OAM
4. Uses multi part ZIPs based on the random number generated from time stamp,size of data, ascii values at specific indices chosen inside the event data
5. Assemble the ZIP with a strategy based on permutation sequence
	The permutation sequence is stored in subscriber profile for configured time
	and based on acknowledgement from user it will be moved to UDSF with documentation.

**Algorithm A2.**
Geo-Hash Generation:
There is assembler function F (D1, D2, D3, metadata) that uses the below input.
1. The public identity
2. permutation sequence
3. random parts of multi part ZIP 

to generate the current Geo-Hash.

the GEOHASH generated can be realised in as   SubjectAltName ::= GeneralNames ex as otherName or registeredID etc., modeled inside UDR or UDSF.
https://datatracker.ietf.org/doc/html/rfc2459#section-4.2.1.7
  

  

