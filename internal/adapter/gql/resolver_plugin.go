package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) Plugin() PluginResolver {
	return &pluginResolver{r}
}

func (r *Resolver) PluginExtension() PluginExtensionResolver {
	return &pluginExtensionResolver{r}
}

type pluginResolver struct{ *Resolver }

func (r *pluginResolver) PropertySchema(ctx context.Context, obj *gqlmodel.Plugin) (*gqlmodel.PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	if obj.PropertySchemaID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).PropertySchema.Load(*obj.PropertySchemaID)
}

func (r *pluginResolver) Scene(ctx context.Context, obj *gqlmodel.Plugin) (*gqlmodel.Scene, error) {
	exit := trace(ctx)
	defer exit()

	if obj.SceneID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Scene.Load(id.SceneID(*obj.SceneID))
}

func (r *pluginResolver) ScenePlugin(ctx context.Context, obj *gqlmodel.Plugin, sceneID *id.ID) (*gqlmodel.ScenePlugin, error) {
	exit := trace(ctx)
	defer exit()

	if sceneID == nil && obj.SceneID != nil {
		sceneID = obj.SceneID
	}
	if sceneID == nil {
		return nil, nil
	}
	s, err := DataLoadersFromContext(ctx).Scene.Load(id.SceneID(*sceneID))
	return s.Plugin(obj.ID), err
}

func (r *pluginResolver) TranslatedName(ctx context.Context, obj *gqlmodel.Plugin, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedName[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Name, nil
}

func (r *pluginResolver) TranslatedDescription(ctx context.Context, obj *gqlmodel.Plugin, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedDescription[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Description, nil
}

type pluginExtensionResolver struct{ *Resolver }

func (r *pluginExtensionResolver) Plugin(ctx context.Context, obj *gqlmodel.PluginExtension) (*gqlmodel.Plugin, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Plugin.Load(obj.PluginID)
}

func (r *pluginExtensionResolver) PropertySchema(ctx context.Context, obj *gqlmodel.PluginExtension) (*gqlmodel.PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.PropertySchemaID)
}

func (r *pluginExtensionResolver) TranslatedName(ctx context.Context, obj *gqlmodel.PluginExtension, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedName[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Name, nil
}

func (r *pluginExtensionResolver) SceneWidget(ctx context.Context, obj *gqlmodel.PluginExtension, sceneID id.ID) (*gqlmodel.SceneWidget, error) {
	exit := trace(ctx)
	defer exit()

	s, err := DataLoadersFromContext(ctx).Scene.Load(id.SceneID(sceneID))
	return s.Widget(obj.PluginID, obj.ExtensionID), err
}

func (r *pluginExtensionResolver) TranslatedDescription(ctx context.Context, obj *gqlmodel.PluginExtension, lang *string) (string, error) {
	exit := trace(ctx)
	defer exit()

	if s, ok := obj.AllTranslatedDescription[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Description, nil
}