// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CertifyVEXStatement_id(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_subject(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_subject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Subject, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.PackageOrArtifact)
	fc.Result = res
	return ec.marshalNPackageOrArtifact2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrArtifact(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_subject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type PackageOrArtifact does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_vulnerability(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_vulnerability(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Vulnerability, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.Vulnerability)
	fc.Result = res
	return ec.marshalNVulnerability2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerability(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_vulnerability(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Vulnerability does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_status(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_status(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Status, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.VexStatus)
	fc.Result = res
	return ec.marshalNVexStatus2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_status(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type VexStatus does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_vexJustification(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_vexJustification(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.VexJustification, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.VexJustification)
	fc.Result = res
	return ec.marshalNVexJustification2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_vexJustification(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type VexJustification does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_statement(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_statement(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Statement, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_statement(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_statusNotes(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_statusNotes(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StatusNotes, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_statusNotes(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_knownSince(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_knownSince(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.KnownSince, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_knownSince(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_origin(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVEXStatement_collector(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVEXStatement) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVEXStatement_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVEXStatement_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVEXStatement",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCertifyVEXStatementSpec(ctx context.Context, obj interface{}) (model.CertifyVEXStatementSpec, error) {
	var it model.CertifyVEXStatementSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "subject", "vulnerability", "status", "vexJustification", "statement", "statusNotes", "knownSince", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "subject":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("subject"))
			it.Subject, err = ec.unmarshalOPackageOrArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "vulnerability":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("vulnerability"))
			it.Vulnerability, err = ec.unmarshalOVulnerabilitySpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilitySpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "status":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("status"))
			it.Status, err = ec.unmarshalOVexStatus2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx, v)
			if err != nil {
				return it, err
			}
		case "vexJustification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("vexJustification"))
			it.VexJustification, err = ec.unmarshalOVexJustification2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx, v)
			if err != nil {
				return it, err
			}
		case "statement":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("statement"))
			it.Statement, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "statusNotes":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("statusNotes"))
			it.StatusNotes, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "knownSince":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("knownSince"))
			it.KnownSince, err = ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageOrArtifactInput(ctx context.Context, obj interface{}) (model.PackageOrArtifactInput, error) {
	var it model.PackageOrArtifactInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "artifact"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			it.Package, err = ec.unmarshalOPkgInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "artifact":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			it.Artifact, err = ec.unmarshalOArtifactInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageOrArtifactSpec(ctx context.Context, obj interface{}) (model.PackageOrArtifactSpec, error) {
	var it model.PackageOrArtifactSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "artifact"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			it.Package, err = ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "artifact":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			it.Artifact, err = ec.unmarshalOArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVexStatementInputSpec(ctx context.Context, obj interface{}) (model.VexStatementInputSpec, error) {
	var it model.VexStatementInputSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"status", "vexJustification", "statement", "statusNotes", "knownSince", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "status":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("status"))
			it.Status, err = ec.unmarshalNVexStatus2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx, v)
			if err != nil {
				return it, err
			}
		case "vexJustification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("vexJustification"))
			it.VexJustification, err = ec.unmarshalNVexJustification2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx, v)
			if err != nil {
				return it, err
			}
		case "statement":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("statement"))
			it.Statement, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "statusNotes":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("statusNotes"))
			it.StatusNotes, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "knownSince":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("knownSince"))
			it.KnownSince, err = ec.unmarshalNTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _PackageOrArtifact(ctx context.Context, sel ast.SelectionSet, obj model.PackageOrArtifact) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.Package:
		return ec._Package(ctx, sel, &obj)
	case *model.Package:
		if obj == nil {
			return graphql.Null
		}
		return ec._Package(ctx, sel, obj)
	case model.Artifact:
		return ec._Artifact(ctx, sel, &obj)
	case *model.Artifact:
		if obj == nil {
			return graphql.Null
		}
		return ec._Artifact(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var certifyVEXStatementImplementors = []string{"CertifyVEXStatement", "Node"}

func (ec *executionContext) _CertifyVEXStatement(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyVEXStatement) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyVEXStatementImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyVEXStatement")
		case "id":

			out.Values[i] = ec._CertifyVEXStatement_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "subject":

			out.Values[i] = ec._CertifyVEXStatement_subject(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "vulnerability":

			out.Values[i] = ec._CertifyVEXStatement_vulnerability(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "status":

			out.Values[i] = ec._CertifyVEXStatement_status(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "vexJustification":

			out.Values[i] = ec._CertifyVEXStatement_vexJustification(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "statement":

			out.Values[i] = ec._CertifyVEXStatement_statement(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "statusNotes":

			out.Values[i] = ec._CertifyVEXStatement_statusNotes(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "knownSince":

			out.Values[i] = ec._CertifyVEXStatement_knownSince(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "origin":

			out.Values[i] = ec._CertifyVEXStatement_origin(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "collector":

			out.Values[i] = ec._CertifyVEXStatement_collector(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNCertifyVEXStatement2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVEXStatement(ctx context.Context, sel ast.SelectionSet, v model.CertifyVEXStatement) graphql.Marshaler {
	return ec._CertifyVEXStatement(ctx, sel, &v)
}

func (ec *executionContext) marshalNCertifyVEXStatement2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVEXStatementᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyVEXStatement) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCertifyVEXStatement2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVEXStatement(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNCertifyVEXStatement2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVEXStatement(ctx context.Context, sel ast.SelectionSet, v *model.CertifyVEXStatement) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyVEXStatement(ctx, sel, v)
}

func (ec *executionContext) marshalNPackageOrArtifact2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrArtifact(ctx context.Context, sel ast.SelectionSet, v model.PackageOrArtifact) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PackageOrArtifact(ctx, sel, v)
}

func (ec *executionContext) unmarshalNPackageOrArtifactInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrArtifactInput(ctx context.Context, v interface{}) (model.PackageOrArtifactInput, error) {
	res, err := ec.unmarshalInputPackageOrArtifactInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNVexJustification2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx context.Context, v interface{}) (model.VexJustification, error) {
	var res model.VexJustification
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNVexJustification2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx context.Context, sel ast.SelectionSet, v model.VexJustification) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNVexStatementInputSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatementInputSpec(ctx context.Context, v interface{}) (model.VexStatementInputSpec, error) {
	res, err := ec.unmarshalInputVexStatementInputSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNVexStatus2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx context.Context, v interface{}) (model.VexStatus, error) {
	var res model.VexStatus
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNVexStatus2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx context.Context, sel ast.SelectionSet, v model.VexStatus) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOCertifyVEXStatementSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVEXStatementSpec(ctx context.Context, v interface{}) (*model.CertifyVEXStatementSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputCertifyVEXStatementSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOPackageOrArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrArtifactSpec(ctx context.Context, v interface{}) (*model.PackageOrArtifactSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputPackageOrArtifactSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOVexJustification2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx context.Context, v interface{}) (*model.VexJustification, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(model.VexJustification)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOVexJustification2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexJustification(ctx context.Context, sel ast.SelectionSet, v *model.VexJustification) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) unmarshalOVexStatus2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx context.Context, v interface{}) (*model.VexStatus, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(model.VexStatus)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOVexStatus2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVexStatus(ctx context.Context, sel ast.SelectionSet, v *model.VexStatus) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

// endregion ***************************** type.gotpl *****************************
