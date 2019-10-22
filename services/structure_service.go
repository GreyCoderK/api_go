package services

import (
	"../dtos"
	. "../model"
	. "../repositories"
)

func CreateStructure(m *Structure, r StructureRepository) dtos.Response {
	operationResult := r.Save(m)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Structure)

	return dtos.Response{Success: true, Data: data}
}

func FindAllStructures(r StructureRepository) dtos.Response {
	operationResult := r.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*Structure)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneStructureById(id uint, r StructureRepository) dtos.Response {
	operationResult := r.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*Structure)

	return dtos.Response{Success: true, Data: data}
}

func UpdateStructureById(id uint, m Structure, r StructureRepository) dtos.Response {
	existingStructureResponse := FindOneStructureById(id, r)

	if !existingStructureResponse.Success {
		return existingStructureResponse
	}

	existingStructure := existingStructureResponse.Data.(*Structure)

	existingStructure.RaisonSocial = m.RaisonSocial

	operationResult := r.Save(existingStructure)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneStructureById(id uint, r StructureRepository) dtos.Response {
	operationResult := r.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteStructureByIds(multiId *dtos.MultiID, r StructureRepository) dtos.Response {
	operationResult := r.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
