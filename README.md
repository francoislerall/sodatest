# Soda test

Test if the UUID of the datasets of SODA database are matching the profiles in the EPD editor.

```bash
go run main.go model.go
```

## Results:

```bash
profile: EN_15804_A2
All the unit groups of the profile indicators has been found (37).
The following indicators could not be found (18):
 - Components for re-use (CRU)
 - Exported electrical energy (EEE)
 - Exported thermal energy (EET)
 - Hazardous waste disposed (HWD)
 - Materials for energy recovery (MER)
 - Materials for recycling (MFR)
 - Use of net fresh water (FW)
 - Non-hazardous waste disposed (NHWD)
 - Use of non renewable primary energy (PENRE)
 - Use of non renewable primary energy resources used as raw materials (PENRM)
 - Use of non renewable secondary fuels (NRSF)
 - Radioactive waste disposed (RWD)
 - Use of renewable primary energy (PERE)
 - Use of renewable primary energy resources used as raw materials (PERM)
 - Use of renewable secondary fuels (RSF)
 - Input of secondary material (SM)
 - Total use of non-renewable primary energy resources (PENRT)
 - Total use of renewable primary energy resources (PERT)%    

profile: EN_15804
All the unit groups of the profile indicators has been found (25).
The following indicators could not be found (18):
 - Components for re-use
 - Exported electrical energy
 - Exported thermal energy
 - Hazardous waste disposed
 - Materials for energy recovery
 - Materials for recycling
 - Use of net fresh water
 - Non hazardous waste dispose
 - Use of non renewable primary energy
 - Use of non renewable primary energy resources used as raw materials
 - Use of non renewable secondary fuels
 - Radioactive waste disposed
 - Use of renewable primary energy
 - Use of renewable primary energy resources used as raw materials
 - Use of renewable secondary fuels
 - Use of secondary material
 - Total use of non renewable primary energy resource
 - Total use of renewable primary energy resources
```
